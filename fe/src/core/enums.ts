export enum Entity {
  LAPS = 'LAPS',
  CARS = 'CARS',
  TRACKS = 'TRACKS',
  GAMES = 'GAMES',
  GEAR = 'GEAR',
}

export type ResourceKey =
  | 'lap'
  | 'car'
  | 'track'
  | 'game'
  | 'wheel'
  | 'pedals'
  | 'cockpit';

export type GearResourceKey = 'wheel' | 'pedals' | 'cockpit';

export type ViewMode = 'list' | 'create' | 'detail' | 'edit';

export type SortOrder = 'asc' | 'desc';

export function isGearResource(resourceKey: ResourceKey): resourceKey is GearResourceKey {
  return (
    resourceKey === 'wheel' ||
    resourceKey === 'pedals' ||
    resourceKey === 'cockpit'
  );
}
