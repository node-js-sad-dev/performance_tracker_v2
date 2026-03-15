import { parseLapTime, RESOURCE_CONFIG } from './consts.ts';
import type {
  ApiEnvelope,
  ListQuery,
  PlainRecord,
  ResourceFormState,
  ResourceListItemMap,
  ResourceListResponseMap,
  ResourceListResult,
  ResourceRecordMap,
  SelectOption,
} from './interfaces.ts';
import type { ResourceKey } from './enums.ts';

function constructUrl(route: string) {
  const baseUrl = import.meta.env.VITE_API_BASE_URL || '/api/v1';
  return `${baseUrl}${route}`;
}

async function request<T>(route: string, options: RequestInit = {}): Promise<T> {
  const response = await fetch(constructUrl(route), options);

  let payload: ApiEnvelope<T> | null = null;
  try {
    payload = (await response.json()) as ApiEnvelope<T>;
  } catch {
    payload = null;
  }

  if (!response.ok || !payload?.success) {
    throw new Error(payload?.error ?? `Request failed with status ${response.status}`);
  }

  return payload.data;
}

function buildQueryString(resourceKey: ResourceKey, query: ListQuery): string {
  const config = RESOURCE_CONFIG[resourceKey];
  const params = new URLSearchParams({
    page: String(query.page),
    limit: String(query.limit),
    sortBy: query.sortBy,
    sortOrder: query.sortOrder,
  });

  if (config.searchKey && query.search?.trim()) {
    params.set(config.searchKey, query.search.trim());
  }

  return params.toString();
}

function normalizeListResponse<K extends ResourceKey>(
  resourceKey: K,
  data: ResourceListResponseMap[K]
): ResourceListResult<ResourceListItemMap[K]> {
  switch (resourceKey) {
    case 'lap': {
      const typedData = data as ResourceListResponseMap['lap'];
      return {
        items: typedData.laps as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
    case 'car': {
      const typedData = data as ResourceListResponseMap['car'];
      return {
        items: typedData.cars as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
    case 'track': {
      const typedData = data as ResourceListResponseMap['track'];
      return {
        items: typedData.tracks as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
    case 'game': {
      const typedData = data as ResourceListResponseMap['game'];
      return {
        items: typedData.games as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
    case 'wheel': {
      const typedData = data as ResourceListResponseMap['wheel'];
      return {
        items: typedData.wheels as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
    case 'pedals': {
      const typedData = data as ResourceListResponseMap['pedals'];
      return {
        items: typedData.pedals as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
    case 'cockpit': {
      const typedData = data as ResourceListResponseMap['cockpit'];
      return {
        items: typedData.cockpits as ResourceListItemMap[K][],
        totalCount: typedData.total_count,
      };
    }
  }
}

function appendTextValue(formData: FormData, key: string, value: unknown) {
  if (typeof value === 'string') {
    formData.append(key, value);
  }
}

function appendFileValue(formData: FormData, key: string, value: unknown) {
  if (value instanceof File) {
    formData.append(key, value);
  }
}

function buildMutationRequest(resourceKey: ResourceKey, formState: ResourceFormState): RequestInit {
  switch (resourceKey) {
    case 'lap':
      return {
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          car_id: Number(formState.car_id),
          track_id: Number(formState.track_id),
          game_id: Number(formState.game_id),
          wheel_id: Number(formState.wheel_id),
          cockpit_id: Number(formState.cockpit_id),
          pedals_id: Number(formState.pedals_id),
          time: parseLapTime(String(formState.time ?? '')),
          is_clear: formState.is_clear === true,
          has_significant_errors: formState.has_significant_errors === true,
        }),
      };
    case 'car':
    case 'track': {
      const formData = new FormData();
      appendTextValue(formData, 'name', formState.name);
      appendTextValue(formData, 'description', formState.description);
      appendFileValue(formData, 'file', formState.imageFile);

      return {
        body: formData,
      };
    }
    case 'game': {
      const formData = new FormData();
      appendTextValue(formData, 'name', formState.name);
      appendFileValue(formData, 'file', formState.imageFile);

      return {
        body: formData,
      };
    }
    case 'wheel':
    case 'pedals':
    case 'cockpit':
      return {
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: String(formState.name ?? ''),
          is_default: formState.is_default === true,
        }),
      };
  }
}

export async function getResourceList<K extends ResourceKey>(
  resourceKey: K,
  query: ListQuery
): Promise<ResourceListResult<ResourceListItemMap[K]>> {
  const config = RESOURCE_CONFIG[resourceKey];
  const queryString = buildQueryString(resourceKey, query);
  const data = await request<ResourceListResponseMap[K]>(
    `${config.apiPath}?${queryString}`
  );

  return normalizeListResponse(resourceKey, data);
}

export async function getResourceById<K extends ResourceKey>(
  resourceKey: K,
  id: number | string
): Promise<ResourceRecordMap[K]> {
  const config = RESOURCE_CONFIG[resourceKey];
  return request<ResourceRecordMap[K]>(`${config.apiPath}/${id}`);
}

export async function createResource<K extends ResourceKey>(
  resourceKey: K,
  formState: ResourceFormState
): Promise<ResourceRecordMap[K]> {
  const config = RESOURCE_CONFIG[resourceKey];
  return request<ResourceRecordMap[K]>(config.apiPath, {
    method: 'POST',
    ...buildMutationRequest(resourceKey, formState),
  });
}

export async function updateResource(
  resourceKey: ResourceKey,
  id: number | string,
  formState: ResourceFormState
): Promise<void> {
  const config = RESOURCE_CONFIG[resourceKey];
  await request(`${config.apiPath}/${id}`, {
    method: 'PATCH',
    ...buildMutationRequest(resourceKey, formState),
  });
}

export async function deleteResource(
  resourceKey: ResourceKey,
  id: number | string
): Promise<void> {
  const config = RESOURCE_CONFIG[resourceKey];
  await request(`${config.apiPath}/${id}`, {
    method: 'DELETE',
  });
}

export async function getSelectOptions(resourceKey: ResourceKey): Promise<SelectOption[]> {
  const config = RESOURCE_CONFIG[resourceKey];
  const list = await getResourceList(resourceKey, {
    page: 1,
    limit: 200,
    sortBy: 'name',
    sortOrder: 'asc',
  });

  return list.items.map((item) => {
    const record = item as unknown as PlainRecord;
    const label = String(record.name ?? `#${record.id ?? '?'}`);

    return {
      label,
      value: String(record.id ?? ''),
      description:
        typeof record.description === 'string' && record.description.trim() !== ''
          ? record.description
          : `Select ${label.toLowerCase()} from ${config.labelPlural.toLowerCase()}.`,
    };
  });
}

