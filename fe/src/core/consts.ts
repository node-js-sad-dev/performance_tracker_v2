import {
  Entity,
  isGearResource,
  type GearResourceKey,
  type ResourceKey,
  type ViewMode,
} from './enums.ts';
import type {
  DetailFieldDefinition,
  FormFieldDefinition,
  LookupCollection,
  PlainRecord,
  ResourceConfig,
  ResourceFormState,
  SelectOption,
  Tone,
} from './interfaces.ts';

function asText(value: unknown, fallback = 'Not set'): string {
  if (value === null || value === undefined) {
    return fallback;
  }

  const stringValue = String(value).trim();
  return stringValue === '' ? fallback : stringValue;
}

function formatDateTime(value: unknown): string {
  if (typeof value !== 'string' || value.trim() === '') {
    return 'Not set';
  }

  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return value;
  }

  return new Intl.DateTimeFormat(undefined, {
    dateStyle: 'medium',
    timeStyle: 'short',
  }).format(date);
}

export function formatLapTime(seconds: number): string {
  const safeMilliseconds = Math.max(0, Math.round(seconds * 1000));
  const minutes = Math.floor(safeMilliseconds / 60000);
  const remainingMilliseconds = safeMilliseconds % 60000;
  const wholeSeconds = Math.floor(remainingMilliseconds / 1000);
  const milliseconds = remainingMilliseconds % 1000;

  return `${minutes}:${String(wholeSeconds).padStart(2, '0')}.${String(
    milliseconds
  ).padStart(3, '0')}`;
}

export function parseLapTime(value: string): number {
  const normalized = value.trim();
  if (normalized === '') {
    throw new Error('Lap time is required.');
  }

  if (/^\d+(\.\d+)?$/.test(normalized)) {
    const seconds = Number(normalized);
    if (Number.isNaN(seconds)) {
      throw new Error('Lap time must be a valid number.');
    }

    return seconds;
  }

  const match = normalized.match(/^(\d+):([0-5]?\d)(?:\.(\d{1,3}))?$/);
  if (!match) {
    throw new Error('Use seconds or m:ss.mmm format for lap time.');
  }

  const minutes = Number(match[1]);
  const seconds = Number(match[2]);
  const milliseconds = Number((match[3] ?? '').padEnd(3, '0') || '0');

  return minutes * 60 + seconds + milliseconds / 1000;
}

function formatBoolean(value: unknown): string {
  return value === true ? 'Yes' : 'No';
}

function formatClearStatus(value: unknown): string {
  return value === true ? 'Clean' : 'Flagged';
}

function formatDefaultStatus(value: unknown): string {
  return value === true ? 'Default' : 'Alt';
}

function resolveOptionLabel(
  options: SelectOption[] | undefined,
  value: unknown,
  fallbackPrefix: string
): string {
  if (typeof value !== 'number') {
    return 'Not set';
  }

  const option = options?.find((item) => item.value === String(value));
  return option?.label ?? `${fallbackPrefix} #${value}`;
}

function toneFromBoolean(value: unknown): Tone {
  return value === true ? 'success' : 'danger';
}

function toneFromDefault(value: unknown): Tone {
  return value === true ? 'accent' : 'muted';
}

function buildNameFields(descriptionHelp?: string): FormFieldDefinition[] {
  return [
    {
      key: 'name',
      label: 'Name',
      type: 'text',
      placeholder: 'Enter a clear display name',
      required: true,
    },
    {
      key: 'description',
      label: 'Description',
      type: 'textarea',
      placeholder: 'Add notes, setup details, or quick reminders',
      help: descriptionHelp,
    },
    {
      key: 'imageFile',
      label: 'Image',
      type: 'file',
      accept: 'image/*',
      help: 'Upload a fresh image to attach or replace the current one.',
    },
  ];
}

function buildGearFields(label: string): FormFieldDefinition[] {
  return [
    {
      key: 'name',
      label: `${label} name`,
      type: 'text',
      placeholder: `Enter ${label.toLowerCase()} name`,
      required: true,
    },
    {
      key: 'is_default',
      label: 'Set as default',
      type: 'checkbox',
      help: `Mark this ${label.toLowerCase()} as the default setup.`,
    },
  ];
}

const createdField: DetailFieldDefinition = {
  label: 'Created',
  getValue: (item) => formatDateTime(item.created_at),
};

export const NAV_ITEMS = [
  {
    label: 'Laps',
    icon: 'LP',
    entity: Entity.LAPS,
    to: '/laps',
  },
  {
    label: 'Cars',
    icon: 'CR',
    entity: Entity.CARS,
    to: '/cars',
  },
  {
    label: 'Tracks',
    icon: 'TR',
    entity: Entity.TRACKS,
    to: '/tracks',
  },
  {
    label: 'Games',
    icon: 'GM',
    entity: Entity.GAMES,
    to: '/games',
  },
  {
    label: 'Gear',
    icon: 'GR',
    entity: Entity.GEAR,
    to: '/gear',
  },
] as const;

export const GEAR_TABS: Array<{
  key: GearResourceKey;
  label: string;
}> = [
  { key: 'wheel', label: 'Wheels' },
  { key: 'pedals', label: 'Pedals' },
  { key: 'cockpit', label: 'Cockpits' },
];

export const RESOURCE_CONFIG: Record<ResourceKey, ResourceConfig> = {
  lap: {
    key: 'lap',
    navEntity: Entity.LAPS,
    slug: 'laps',
    apiPath: '/lap',
    labelSingular: 'Lap',
    labelPlural: 'Laps',
    headerTitle: 'Laps Telemetry',
    createActionLabel: '+ Log New Lap',
    emptyStateLabel: 'No laps found yet. Start by logging a new lap.',
    defaultSortBy: 'date',
    defaultSortOrder: 'desc',
    tableColumns: [
      {
        key: 'date',
        label: 'Date',
        sortBy: 'date',
        kind: 'datetime',
        getValue: (item) => asText(item.date),
      },
      {
        key: 'game',
        label: 'Game',
        sortBy: 'game',
        getValue: (item) => asText(item.game),
      },
      {
        key: 'car',
        label: 'Car',
        sortBy: 'car',
        getValue: (item) => asText(item.car),
      },
      {
        key: 'track',
        label: 'Track',
        sortBy: 'track',
        getValue: (item) => asText(item.track),
      },
      {
        key: 'time',
        label: 'Time',
        sortBy: 'time',
        align: 'right',
        kind: 'time',
        getValue: (item) => asText(item.time),
      },
      {
        key: 'clear',
        label: 'Clear',
        sortBy: 'clear',
        align: 'center',
        kind: 'badge',
        getValue: (item) => formatClearStatus(item.clear),
        getTone: (item) => toneFromBoolean(item.clear),
      },
    ],
    formFields: [
      {
        key: 'car_id',
        label: 'Car',
        type: 'select',
        optionResource: 'car',
        required: true,
      },
      {
        key: 'track_id',
        label: 'Track',
        type: 'select',
        optionResource: 'track',
        required: true,
      },
      {
        key: 'game_id',
        label: 'Game',
        type: 'select',
        optionResource: 'game',
        required: true,
      },
      {
        key: 'wheel_id',
        label: 'Wheel',
        type: 'select',
        optionResource: 'wheel',
        required: true,
      },
      {
        key: 'pedals_id',
        label: 'Pedals',
        type: 'select',
        optionResource: 'pedals',
        required: true,
      },
      {
        key: 'cockpit_id',
        label: 'Cockpit',
        type: 'select',
        optionResource: 'cockpit',
        required: true,
      },
      {
        key: 'time',
        label: 'Lap time',
        type: 'lap-time',
        placeholder: '1:54.302 or 114.302',
        required: true,
        help: 'Use lap format or total seconds.',
      },
      {
        key: 'is_clear',
        label: 'Clean lap',
        type: 'checkbox',
      },
      {
        key: 'has_significant_errors',
        label: 'Significant errors',
        type: 'checkbox',
      },
    ],
    detailFields: [
      {
        label: 'Car',
        getValue: (item, lookups) =>
          resolveOptionLabel(lookups.car, item.car_id, 'Car'),
      },
      {
        label: 'Track',
        getValue: (item, lookups) =>
          resolveOptionLabel(lookups.track, item.track_id, 'Track'),
      },
      {
        label: 'Game',
        getValue: (item, lookups) =>
          resolveOptionLabel(lookups.game, item.game_id, 'Game'),
      },
      {
        label: 'Wheel',
        getValue: (item, lookups) =>
          resolveOptionLabel(lookups.wheel, item.wheel_id, 'Wheel'),
      },
      {
        label: 'Pedals',
        getValue: (item, lookups) =>
          resolveOptionLabel(lookups.pedals, item.pedals_id, 'Pedals'),
      },
      {
        label: 'Cockpit',
        getValue: (item, lookups) =>
          resolveOptionLabel(lookups.cockpit, item.cockpit_id, 'Cockpit'),
      },
      {
        label: 'Lap time',
        getValue: (item) => formatLapTime(Number(item.time ?? 0)),
        tone: 'accent',
      },
      {
        label: 'Clean lap',
        getValue: (item) => formatClearStatus(item.is_clear),
        tone: (item) => toneFromBoolean(item.is_clear),
      },
      {
        label: 'Significant errors',
        getValue: (item) => formatBoolean(item.has_significant_errors),
        tone: (item) => toneFromBoolean(!(item.has_significant_errors as boolean)),
      },
      createdField,
    ],
  },
  car: {
    key: 'car',
    navEntity: Entity.CARS,
    slug: 'cars',
    apiPath: '/car',
    labelSingular: 'Car',
    labelPlural: 'Cars',
    headerTitle: 'Cars Telemetry',
    createActionLabel: '+ Add New Car',
    emptyStateLabel: 'No cars found yet. Add your first car to get started.',
    defaultSortBy: 'created_at',
    defaultSortOrder: 'desc',
    searchKey: 'name',
    searchPlaceholder: 'Search cars by name',
    tableColumns: [
      {
        key: 'name',
        label: 'Name',
        sortBy: 'name',
        getValue: (item) => asText(item.name),
      },
      {
        key: 'description',
        label: 'Description',
        getValue: (item) => asText(item.description, 'No description'),
      },
      {
        key: 'created_at',
        label: 'Created',
        sortBy: 'created_at',
        kind: 'datetime',
        getValue: (item) => formatDateTime(item.created_at),
      },
    ],
    formFields: buildNameFields('Store setup notes, tire choice, or handling reminders.'),
    detailFields: [
      {
        label: 'Name',
        getValue: (item) => asText(item.name),
      },
      {
        label: 'Description',
        getValue: (item) => asText(item.description, 'No description'),
      },
      createdField,
    ],
  },
  track: {
    key: 'track',
    navEntity: Entity.TRACKS,
    slug: 'tracks',
    apiPath: '/track',
    labelSingular: 'Track',
    labelPlural: 'Tracks',
    headerTitle: 'Tracks Telemetry',
    createActionLabel: '+ Add New Track',
    emptyStateLabel: 'No tracks found yet. Add a track to populate the grid.',
    defaultSortBy: 'created_at',
    defaultSortOrder: 'desc',
    searchKey: 'name',
    searchPlaceholder: 'Search tracks by name',
    tableColumns: [
      {
        key: 'name',
        label: 'Name',
        sortBy: 'name',
        getValue: (item) => asText(item.name),
      },
      {
        key: 'description',
        label: 'Description',
        getValue: (item) => asText(item.description, 'No description'),
      },
      {
        key: 'created_at',
        label: 'Created',
        sortBy: 'created_at',
        kind: 'datetime',
        getValue: (item) => formatDateTime(item.created_at),
      },
    ],
    formFields: buildNameFields('Track notes work well here: braking markers, curbs, or setup comments.'),
    detailFields: [
      {
        label: 'Name',
        getValue: (item) => asText(item.name),
      },
      {
        label: 'Description',
        getValue: (item) => asText(item.description, 'No description'),
      },
      createdField,
    ],
  },
  game: {
    key: 'game',
    navEntity: Entity.GAMES,
    slug: 'games',
    apiPath: '/game',
    labelSingular: 'Game',
    labelPlural: 'Games',
    headerTitle: 'Games Telemetry',
    createActionLabel: '+ Add New Game',
    emptyStateLabel: 'No games found yet. Add the sims you want to track.',
    defaultSortBy: 'created_at',
    defaultSortOrder: 'desc',
    searchKey: 'name',
    searchPlaceholder: 'Search games by name',
    tableColumns: [
      {
        key: 'name',
        label: 'Name',
        sortBy: 'name',
        getValue: (item) => asText(item.name),
      },
      {
        key: 'created_at',
        label: 'Created',
        sortBy: 'created_at',
        kind: 'datetime',
        getValue: (item) => formatDateTime(item.created_at),
      },
    ],
    formFields: [
      {
        key: 'name',
        label: 'Name',
        type: 'text',
        placeholder: 'Enter the simulator name',
        required: true,
      },
      {
        key: 'imageFile',
        label: 'Cover image',
        type: 'file',
        accept: 'image/*',
        requiredOnCreate: true,
        help: 'The backend currently requires an image when creating a game.',
      },
    ],
    detailFields: [
      {
        label: 'Name',
        getValue: (item) => asText(item.name),
      },
      createdField,
    ],
  },
  wheel: {
    key: 'wheel',
    navEntity: Entity.GEAR,
    slug: 'wheels',
    apiPath: '/wheel',
    labelSingular: 'Wheel',
    labelPlural: 'Wheels',
    headerTitle: 'Gear Telemetry',
    createActionLabel: '+ Add New Wheel',
    emptyStateLabel: 'No wheels found yet. Add the wheelbases or rims you use.',
    defaultSortBy: 'created_at',
    defaultSortOrder: 'desc',
    searchKey: 'name',
    searchPlaceholder: 'Search wheels by name',
    tableColumns: [
      {
        key: 'name',
        label: 'Name',
        sortBy: 'name',
        getValue: (item) => asText(item.name),
      },
      {
        key: 'is_default',
        label: 'Default',
        sortBy: 'is_default',
        align: 'center',
        kind: 'badge',
        getValue: (item) => formatDefaultStatus(item.is_default),
        getTone: (item) => toneFromDefault(item.is_default),
      },
      {
        key: 'created_at',
        label: 'Created',
        sortBy: 'created_at',
        kind: 'datetime',
        getValue: (item) => formatDateTime(item.created_at),
      },
    ],
    formFields: buildGearFields('Wheel'),
    detailFields: [
      {
        label: 'Name',
        getValue: (item) => asText(item.name),
      },
      {
        label: 'Default setup',
        getValue: (item) => formatDefaultStatus(item.is_default),
        tone: (item) => toneFromDefault(item.is_default),
      },
      createdField,
    ],
  },
  pedals: {
    key: 'pedals',
    navEntity: Entity.GEAR,
    slug: 'pedals',
    apiPath: '/pedals',
    labelSingular: 'Pedals',
    labelPlural: 'Pedals',
    headerTitle: 'Gear Telemetry',
    createActionLabel: '+ Add New Pedals',
    emptyStateLabel: 'No pedal sets found yet. Add the ones in your rotation.',
    defaultSortBy: 'created_at',
    defaultSortOrder: 'desc',
    searchKey: 'name',
    searchPlaceholder: 'Search pedals by name',
    tableColumns: [
      {
        key: 'name',
        label: 'Name',
        sortBy: 'name',
        getValue: (item) => asText(item.name),
      },
      {
        key: 'is_default',
        label: 'Default',
        sortBy: 'is_default',
        align: 'center',
        kind: 'badge',
        getValue: (item) => formatDefaultStatus(item.is_default),
        getTone: (item) => toneFromDefault(item.is_default),
      },
      {
        key: 'created_at',
        label: 'Created',
        sortBy: 'created_at',
        kind: 'datetime',
        getValue: (item) => formatDateTime(item.created_at),
      },
    ],
    formFields: buildGearFields('Pedals'),
    detailFields: [
      {
        label: 'Name',
        getValue: (item) => asText(item.name),
      },
      {
        label: 'Default setup',
        getValue: (item) => formatDefaultStatus(item.is_default),
        tone: (item) => toneFromDefault(item.is_default),
      },
      createdField,
    ],
  },
  cockpit: {
    key: 'cockpit',
    navEntity: Entity.GEAR,
    slug: 'cockpits',
    apiPath: '/cockpit',
    labelSingular: 'Cockpit',
    labelPlural: 'Cockpits',
    headerTitle: 'Gear Telemetry',
    createActionLabel: '+ Add New Cockpit',
    emptyStateLabel: 'No cockpits found yet. Add the rigs you want to track.',
    defaultSortBy: 'created_at',
    defaultSortOrder: 'desc',
    searchKey: 'name',
    searchPlaceholder: 'Search cockpits by name',
    tableColumns: [
      {
        key: 'name',
        label: 'Name',
        sortBy: 'name',
        getValue: (item) => asText(item.name),
      },
      {
        key: 'is_default',
        label: 'Default',
        sortBy: 'is_default',
        align: 'center',
        kind: 'badge',
        getValue: (item) => formatDefaultStatus(item.is_default),
        getTone: (item) => toneFromDefault(item.is_default),
      },
      {
        key: 'created_at',
        label: 'Created',
        sortBy: 'created_at',
        kind: 'datetime',
        getValue: (item) => formatDateTime(item.created_at),
      },
    ],
    formFields: buildGearFields('Cockpit'),
    detailFields: [
      {
        label: 'Name',
        getValue: (item) => asText(item.name),
      },
      {
        label: 'Default setup',
        getValue: (item) => formatDefaultStatus(item.is_default),
        tone: (item) => toneFromDefault(item.is_default),
      },
      createdField,
    ],
  },
};

export function buildResourcePath(
  resourceKey: ResourceKey,
  mode: ViewMode,
  id?: number | string
): string {
  const config = RESOURCE_CONFIG[resourceKey];
  const basePath = isGearResource(resourceKey)
    ? `/gear/${config.slug}`
    : `/${config.slug}`;

  if (mode === 'list') {
    return basePath;
  }

  if (mode === 'create') {
    return `${basePath}/new`;
  }

  if (!id) {
    return basePath;
  }

  if (mode === 'detail') {
    return `${basePath}/${id}`;
  }

  return `${basePath}/${id}/edit`;
}

export function createEmptyFormState(resourceKey: ResourceKey): ResourceFormState {
  switch (resourceKey) {
    case 'lap':
      return {
        car_id: '',
        track_id: '',
        game_id: '',
        wheel_id: '',
        pedals_id: '',
        cockpit_id: '',
        time: '',
        is_clear: true,
        has_significant_errors: false,
      };
    case 'car':
    case 'track':
      return {
        name: '',
        description: '',
        imageFile: null,
      };
    case 'game':
      return {
        name: '',
        imageFile: null,
      };
    case 'wheel':
    case 'pedals':
    case 'cockpit':
      return {
        name: '',
        is_default: false,
      };
  }
}

export function mapRecordToForm(
  resourceKey: ResourceKey,
  record: PlainRecord
): ResourceFormState {
  switch (resourceKey) {
    case 'lap':
      return {
        car_id: String(record.car_id ?? ''),
        track_id: String(record.track_id ?? ''),
        game_id: String(record.game_id ?? ''),
        wheel_id: String(record.wheel_id ?? ''),
        pedals_id: String(record.pedals_id ?? ''),
        cockpit_id: String(record.cockpit_id ?? ''),
        time: formatLapTime(Number(record.time ?? 0)),
        is_clear: record.is_clear === true,
        has_significant_errors: record.has_significant_errors === true,
      };
    case 'car':
    case 'track':
      return {
        name: asText(record.name, ''),
        description: asText(record.description, ''),
        imageFile: null,
      };
    case 'game':
      return {
        name: asText(record.name, ''),
        imageFile: null,
      };
    case 'wheel':
    case 'pedals':
    case 'cockpit':
      return {
        name: asText(record.name, ''),
        is_default: record.is_default === true,
      };
  }
}

export function getResourceHeading(resourceKey: ResourceKey, record: PlainRecord): string {
  if (resourceKey === 'lap') {
    return `Lap #${asText(record.id)}`;
  }

  return asText(record.name, RESOURCE_CONFIG[resourceKey].labelSingular);
}

export function getImageFromRecord(record: PlainRecord): string | null {
  const image = record.image;
  return typeof image === 'string' && image.trim() !== '' ? image : null;
}

export function getDetailEntries(
  resourceKey: ResourceKey,
  record: PlainRecord,
  lookups: LookupCollection
) {
  return RESOURCE_CONFIG[resourceKey].detailFields.map((field) => ({
    label: field.label,
    value: field.getValue(record, lookups),
    tone:
      typeof field.tone === 'function'
        ? field.tone(record)
        : field.tone ?? 'default',
  }));
}

