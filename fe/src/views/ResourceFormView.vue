<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import GearTabs from '../components/GearTabs/GearTabs.vue';
import NewEntryButton from '../components/NewEntryButton/NewEntryButton.vue';
import TopBar from '../components/TopBar/TopBar.vue';
import {
  RESOURCE_CONFIG,
  buildResourcePath,
  createEmptyFormState,
  getImageFromRecord,
  getResourceHeading,
  mapRecordToForm,
  parseLapTime,
} from '../core/consts.ts';
import {
  createResource,
  getResourceById,
  getSelectOptions,
  updateResource,
} from '../core/fetcher.ts';
import {
  Entity,
  type GearResourceKey,
  type ResourceKey,
  type ViewMode,
} from '../core/enums.ts';
import type {
  FormFieldDefinition,
  LookupCollection,
  PlainRecord,
  ResourceFormState,
  ResourceFormValue,
} from '../core/interfaces.ts';

const route = useRoute();
const router = useRouter();

const loading = ref(false);
const saving = ref(false);
const error = ref<string | null>(null);
const formState = ref<ResourceFormState>({});
const record = ref<PlainRecord | null>(null);
const lookups = ref<LookupCollection>({});

const resourceKey = computed(() => route.meta.resourceKey as ResourceKey);
const mode = computed(() => route.meta.mode as ViewMode);
const config = computed(() => RESOURCE_CONFIG[resourceKey.value]);
const id = computed(() => Number(route.params.id ?? 0));
const isGearSection = computed(() => config.value.navEntity === Entity.GEAR);
const gearResourceKey = computed(() => resourceKey.value as GearResourceKey);
const existingImage = computed(() => (record.value ? getImageFromRecord(record.value) : null));
const cancelPath = computed(() =>
  mode.value === 'create'
    ? buildResourcePath(resourceKey.value, 'list')
    : buildResourcePath(resourceKey.value, 'detail', id.value)
);
const pageTitle = computed(() => {
  if (mode.value === 'create') {
    return `Create ${config.value.labelSingular}`;
  }

  if (!record.value) {
    return `Edit ${config.value.labelSingular}`;
  }

  return `Edit ${getResourceHeading(resourceKey.value, record.value)}`;
});

function setField(key: string, value: ResourceFormValue) {
  formState.value = {
    ...formState.value,
    [key]: value,
  };
}

function getTextValue(key: string): string {
  const value = formState.value[key];
  if (typeof value === 'string') {
    return value;
  }

  if (typeof value === 'number') {
    return String(value);
  }

  return '';
}

function getBooleanValue(key: string): boolean {
  return formState.value[key] === true;
}

function getFieldOptions(field: FormFieldDefinition) {
  return field.optionResource ? lookups.value[field.optionResource] ?? [] : [];
}

function isFieldRequired(field: FormFieldDefinition) {
  return field.required || (mode.value === 'create' && field.requiredOnCreate === true);
}

function getFileName(key: string): string {
  const value = formState.value[key];
  return value instanceof File ? value.name : '';
}

function handleTextInput(key: string, event: Event) {
  const element = event.target as HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement;
  setField(key, element.value);
}

function handleCheckboxInput(key: string, event: Event) {
  const element = event.target as HTMLInputElement;
  setField(key, element.checked);
}

function handleFileInput(key: string, event: Event) {
  const element = event.target as HTMLInputElement;
  setField(key, element.files?.[0] ?? null);
}

async function loadLookups() {
  if (resourceKey.value !== 'lap') {
    lookups.value = {};
    return;
  }

  const lookupKeys: ResourceKey[] = ['car', 'track', 'game', 'wheel', 'pedals', 'cockpit'];
  const entries = await Promise.all(
    lookupKeys.map(async (lookupKey) => [lookupKey, await getSelectOptions(lookupKey)] as const)
  );

  lookups.value = Object.fromEntries(entries) as LookupCollection;
}

async function loadView() {
  loading.value = true;
  error.value = null;

  try {
    await loadLookups();

    if (mode.value === 'edit') {
      if (!Number.isInteger(id.value) || id.value <= 0) {
        throw new Error('Invalid record identifier.');
      }

      const entity = await getResourceById(resourceKey.value, id.value);
      record.value = entity as unknown as PlainRecord;
      formState.value = mapRecordToForm(resourceKey.value, entity as unknown as PlainRecord);
      return;
    }

    record.value = null;
    formState.value = createEmptyFormState(resourceKey.value);
  } catch (loadError) {
    error.value = loadError instanceof Error ? loadError.message : 'Failed to load form.';
  } finally {
    loading.value = false;
  }
}

watch([resourceKey, mode, id], () => {
  void loadView();
}, { immediate: true });

function validateForm(): string | null {
  for (const field of config.value.formFields) {
    if (!isFieldRequired(field)) {
      continue;
    }

    const value = formState.value[field.key];

    if (field.type === 'checkbox') {
      continue;
    }

    if (field.type === 'file') {
      if (!(value instanceof File)) {
        return `${field.label} is required.`;
      }
      continue;
    }

    if (typeof value === 'string' && value.trim() === '') {
      return `${field.label} is required.`;
    }

    if (value === null || value === undefined || value === '') {
      return `${field.label} is required.`;
    }
  }

  if (resourceKey.value === 'lap') {
    try {
      parseLapTime(getTextValue('time'));
    } catch (validationError) {
      return validationError instanceof Error ? validationError.message : 'Invalid lap time.';
    }
  }

  return null;
}

async function handleSubmit() {
  const validationError = validateForm();
  if (validationError) {
    error.value = validationError;
    return;
  }

  saving.value = true;
  error.value = null;

  try {
    if (mode.value === 'create') {
      const created = await createResource(resourceKey.value, formState.value);
      const createdRecord = created as unknown as PlainRecord;
      await router.push(buildResourcePath(resourceKey.value, 'detail', Number(createdRecord.id)));
      return;
    }

    await updateResource(resourceKey.value, id.value, formState.value);
    await router.push(buildResourcePath(resourceKey.value, 'detail', id.value));
  } catch (saveError) {
    error.value = saveError instanceof Error ? saveError.message : 'Failed to save data.';
  } finally {
    saving.value = false;
  }
}
</script>

<template>
  <section>
    <TopBar
      :eyebrow="mode === 'create' ? 'Create record' : 'Edit record'"
      :title="pageTitle"
      :subtitle="'The form shape is driven by the backend routes already present in the repo.'">
      <template #actions>
        <NewEntryButton :to="cancelPath" label="Back" variant="secondary" />
      </template>
    </TopBar>

    <GearTabs v-if="isGearSection" :resource-key="gearResourceKey" />

    <div v-if="error" class="status-banner error">
      {{ error }}
    </div>

    <div v-if="loading" class="form-card">
      <div class="loading-state">Loading form...</div>
    </div>

    <form v-else class="form-card" @submit.prevent="handleSubmit">
      <div class="field-grid">
        <div
          v-for="field in config.formFields"
          :key="field.key"
          class="field-group"
          :class="{ 'span-2': field.type === 'textarea' || field.type === 'file' }">
          <label class="field-label" :for="field.key">
            <span>{{ field.label }}</span>
            <span v-if="isFieldRequired(field)" class="tone-accent">Required</span>
          </label>

          <input
            v-if="field.type === 'text' || field.type === 'lap-time' || field.type === 'number'"
            :id="field.key"
            class="field-input"
            :type="field.type === 'number' ? 'number' : 'text'"
            :placeholder="field.placeholder"
            :min="field.min"
            :step="field.step"
            :value="getTextValue(field.key)"
            @input="handleTextInput(field.key, $event)" />

          <textarea
            v-else-if="field.type === 'textarea'"
            :id="field.key"
            class="field-textarea"
            :placeholder="field.placeholder"
            :value="getTextValue(field.key)"
            @input="handleTextInput(field.key, $event)" />

          <select
            v-else-if="field.type === 'select'"
            :id="field.key"
            class="field-select"
            :value="getTextValue(field.key)"
            @change="handleTextInput(field.key, $event)">
            <option value="">Select {{ field.label.toLowerCase() }}</option>
            <option v-for="option in getFieldOptions(field)" :key="option.value" :value="option.value">
              {{ option.label }}
            </option>
          </select>

          <label v-else-if="field.type === 'checkbox'" class="checkbox-row" :for="field.key">
            <input
              :id="field.key"
              type="checkbox"
              :checked="getBooleanValue(field.key)"
              @change="handleCheckboxInput(field.key, $event)" />
            <span>{{ field.label }}</span>
          </label>

          <div v-else-if="field.type === 'file'" class="field-group">
            <img
              v-if="existingImage"
              :src="existingImage"
              :alt="`${config.labelSingular} preview`"
              class="image-preview" />
            <input
              :id="field.key"
              class="field-input"
              type="file"
              :accept="field.accept"
              @change="handleFileInput(field.key, $event)" />
            <p v-if="getFileName(field.key)" class="preview-note">
              Selected file: {{ getFileName(field.key) }}
            </p>
          </div>

          <p v-if="field.help" class="field-help">{{ field.help }}</p>
        </div>
      </div>

      <div class="button-row">
        <NewEntryButton :to="cancelPath" label="Cancel" variant="secondary" />
        <NewEntryButton
          :label="saving ? 'Saving...' : mode === 'create' ? 'Create record' : 'Save changes'"
          type="submit"
          :disabled="saving" />
      </div>
    </form>
  </section>
</template>

