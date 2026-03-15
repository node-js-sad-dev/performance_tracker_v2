import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router';

import ResourceDetailView from '../views/ResourceDetailView.vue';
import ResourceFormView from '../views/ResourceFormView.vue';
import ResourceListView from '../views/ResourceListView.vue';
import { Entity } from '../core/enums.ts';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/laps',
  },
  {
    path: '/laps',
    name: 'lap-list',
    component: ResourceListView,
    meta: { entity: Entity.LAPS, resourceKey: 'lap', mode: 'list' },
  },
  {
    path: '/laps/new',
    name: 'lap-create',
    component: ResourceFormView,
    meta: { entity: Entity.LAPS, resourceKey: 'lap', mode: 'create' },
  },
  {
    path: '/laps/:id',
    name: 'lap-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.LAPS, resourceKey: 'lap', mode: 'detail' },
  },
  {
    path: '/laps/:id/edit',
    name: 'lap-edit',
    component: ResourceFormView,
    meta: { entity: Entity.LAPS, resourceKey: 'lap', mode: 'edit' },
  },
  {
    path: '/cars',
    name: 'car-list',
    component: ResourceListView,
    meta: { entity: Entity.CARS, resourceKey: 'car', mode: 'list' },
  },
  {
    path: '/cars/new',
    name: 'car-create',
    component: ResourceFormView,
    meta: { entity: Entity.CARS, resourceKey: 'car', mode: 'create' },
  },
  {
    path: '/cars/:id',
    name: 'car-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.CARS, resourceKey: 'car', mode: 'detail' },
  },
  {
    path: '/cars/:id/edit',
    name: 'car-edit',
    component: ResourceFormView,
    meta: { entity: Entity.CARS, resourceKey: 'car', mode: 'edit' },
  },
  {
    path: '/tracks',
    name: 'track-list',
    component: ResourceListView,
    meta: { entity: Entity.TRACKS, resourceKey: 'track', mode: 'list' },
  },
  {
    path: '/tracks/new',
    name: 'track-create',
    component: ResourceFormView,
    meta: { entity: Entity.TRACKS, resourceKey: 'track', mode: 'create' },
  },
  {
    path: '/tracks/:id',
    name: 'track-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.TRACKS, resourceKey: 'track', mode: 'detail' },
  },
  {
    path: '/tracks/:id/edit',
    name: 'track-edit',
    component: ResourceFormView,
    meta: { entity: Entity.TRACKS, resourceKey: 'track', mode: 'edit' },
  },
  {
    path: '/games',
    name: 'game-list',
    component: ResourceListView,
    meta: { entity: Entity.GAMES, resourceKey: 'game', mode: 'list' },
  },
  {
    path: '/games/new',
    name: 'game-create',
    component: ResourceFormView,
    meta: { entity: Entity.GAMES, resourceKey: 'game', mode: 'create' },
  },
  {
    path: '/games/:id',
    name: 'game-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.GAMES, resourceKey: 'game', mode: 'detail' },
  },
  {
    path: '/games/:id/edit',
    name: 'game-edit',
    component: ResourceFormView,
    meta: { entity: Entity.GAMES, resourceKey: 'game', mode: 'edit' },
  },
  {
    path: '/gear',
    redirect: '/gear/wheels',
  },
  {
    path: '/gear/wheels',
    name: 'wheel-list',
    component: ResourceListView,
    meta: { entity: Entity.GEAR, resourceKey: 'wheel', mode: 'list' },
  },
  {
    path: '/gear/wheels/new',
    name: 'wheel-create',
    component: ResourceFormView,
    meta: { entity: Entity.GEAR, resourceKey: 'wheel', mode: 'create' },
  },
  {
    path: '/gear/wheels/:id',
    name: 'wheel-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.GEAR, resourceKey: 'wheel', mode: 'detail' },
  },
  {
    path: '/gear/wheels/:id/edit',
    name: 'wheel-edit',
    component: ResourceFormView,
    meta: { entity: Entity.GEAR, resourceKey: 'wheel', mode: 'edit' },
  },
  {
    path: '/gear/pedals',
    name: 'pedals-list',
    component: ResourceListView,
    meta: { entity: Entity.GEAR, resourceKey: 'pedals', mode: 'list' },
  },
  {
    path: '/gear/pedals/new',
    name: 'pedals-create',
    component: ResourceFormView,
    meta: { entity: Entity.GEAR, resourceKey: 'pedals', mode: 'create' },
  },
  {
    path: '/gear/pedals/:id',
    name: 'pedals-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.GEAR, resourceKey: 'pedals', mode: 'detail' },
  },
  {
    path: '/gear/pedals/:id/edit',
    name: 'pedals-edit',
    component: ResourceFormView,
    meta: { entity: Entity.GEAR, resourceKey: 'pedals', mode: 'edit' },
  },
  {
    path: '/gear/cockpits',
    name: 'cockpit-list',
    component: ResourceListView,
    meta: { entity: Entity.GEAR, resourceKey: 'cockpit', mode: 'list' },
  },
  {
    path: '/gear/cockpits/new',
    name: 'cockpit-create',
    component: ResourceFormView,
    meta: { entity: Entity.GEAR, resourceKey: 'cockpit', mode: 'create' },
  },
  {
    path: '/gear/cockpits/:id',
    name: 'cockpit-detail',
    component: ResourceDetailView,
    meta: { entity: Entity.GEAR, resourceKey: 'cockpit', mode: 'detail' },
  },
  {
    path: '/gear/cockpits/:id/edit',
    name: 'cockpit-edit',
    component: ResourceFormView,
    meta: { entity: Entity.GEAR, resourceKey: 'cockpit', mode: 'edit' },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
