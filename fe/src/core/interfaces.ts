import type { Entity, ResourceKey, SortOrder } from './enums.ts';

export interface ApiEnvelope<T> {
  success: boolean;
  data: T;
  error?: string;
}

export interface ListQuery {
  page: number;
  limit: number;
  sortBy: string;
  sortOrder: SortOrder;
  search?: string;
}

export interface ResourceListResult<T> {
  items: T[];
  totalCount: number;
}

export interface SelectOption {
  label: string;
  value: string;
  description?: string;
}

export interface LapListItem {
  id: number;
  date: string;
  game: string;
  car: string;
  track: string;
  time: string;
  clear: boolean;
}

export interface LapRecord {
  id: number;
  car_id: number;
  track_id: number;
  game_id: number;
  wheel_id: number;
  cockpit_id: number;
  pedals_id: number;
  time: number;
  is_clear: boolean;
  has_significant_errors: boolean;
  created_at: string;
}

export interface NamedResourceBase {
  id: number;
  name: string;
  created_at: string;
}

export interface CarRecord extends NamedResourceBase {
  description: string | null;
  image: string | null;
}

export interface TrackRecord extends NamedResourceBase {
  description: string | null;
  image: string | null;
}

export interface GameRecord extends NamedResourceBase {
  image: string | null;
}

export interface WheelRecord extends NamedResourceBase {
  is_default: boolean;
}

export interface PedalsRecord extends NamedResourceBase {
  is_default: boolean;
}

export interface CockpitRecord extends NamedResourceBase {
  is_default: boolean;
}

export interface ResourceRecordMap {
  lap: LapRecord;
  car: CarRecord;
  track: TrackRecord;
  game: GameRecord;
  wheel: WheelRecord;
  pedals: PedalsRecord;
  cockpit: CockpitRecord;
}

export interface ResourceListItemMap {
  lap: LapListItem;
  car: CarRecord;
  track: TrackRecord;
  game: GameRecord;
  wheel: WheelRecord;
  pedals: PedalsRecord;
  cockpit: CockpitRecord;
}

export interface ResourceListResponseMap {
  lap: {
    laps: LapListItem[];
    total_count: number;
  };
  car: {
    cars: CarRecord[];
    total_count: number;
  };
  track: {
    tracks: TrackRecord[];
    total_count: number;
  };
  game: {
    games: GameRecord[];
    total_count: number;
  };
  wheel: {
    wheels: WheelRecord[];
    total_count: number;
  };
  pedals: {
    pedals: PedalsRecord[];
    total_count: number;
  };
  cockpit: {
    cockpits: CockpitRecord[];
    total_count: number;
  };
}

export type ResourceFormValue = string | number | boolean | File | null;
export type ResourceFormState = Record<string, ResourceFormValue>;
export type PlainRecord = Record<string, unknown>;

export type FieldType =
  | 'text'
  | 'textarea'
  | 'number'
  | 'checkbox'
  | 'select'
  | 'file'
  | 'lap-time';

export type Tone = 'default' | 'muted' | 'accent' | 'success' | 'danger';

export interface TableColumnDefinition {
  key: string;
  label: string;
  sortBy?: string;
  align?: 'left' | 'right' | 'center';
  kind?: 'text' | 'time' | 'badge' | 'datetime';
  getValue: (item: PlainRecord) => string | number | boolean | null | undefined;
  getTone?: (item: PlainRecord) => Tone;
}

export interface FormFieldDefinition {
  key: string;
  label: string;
  type: FieldType;
  placeholder?: string;
  help?: string;
  required?: boolean;
  requiredOnCreate?: boolean;
  optionResource?: ResourceKey;
  accept?: string;
  min?: number;
  step?: string;
}

export interface DetailFieldDefinition {
  label: string;
  getValue: (item: PlainRecord, lookups: LookupCollection) => string;
  tone?: Tone | ((item: PlainRecord) => Tone);
}

export interface ResourceConfig {
  key: ResourceKey;
  navEntity: Entity;
  slug: string;
  apiPath: string;
  labelSingular: string;
  labelPlural: string;
  headerTitle: string;
  createActionLabel: string;
  emptyStateLabel: string;
  defaultSortBy: string;
  defaultSortOrder: SortOrder;
  searchKey?: string;
  searchPlaceholder?: string;
  tableColumns: TableColumnDefinition[];
  formFields: FormFieldDefinition[];
  detailFields: DetailFieldDefinition[];
}

export type LookupCollection = Partial<Record<ResourceKey, SelectOption[]>>;
