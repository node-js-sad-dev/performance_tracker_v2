performance tracker app description

1. default page - list of all results. From fe - simple page with table where each record will have its actions - edit, delete, open detailed and so on. Table need to have pagination and sorting. Also, page need to have separate button to add new result (think how implement it, maybe first record of table will always be the add new result form)

2. add new record/edit existing page: just simple page with all fields editable

3. Detailed view page - similar to previous but will have all fields read-only and possibly additional information (for v1 additional information will not be present)

For each entity there will be 3 pages - list, edit/create and detailed. design will be identical for each entity, only difference will be set of fields

4. On each page will be side menu where will be navigation. Full list of pages available from side menu: Laps (default), Cars, Tracks, Games, Gear (combined page where will be wheels, pedals, cockpits, gearboxes and so on)

Database entities:

1. Lap
  id
  createdAt
  carId
  trackId
  gameId
  wheelId
  cockpitId
  pedalsId
  gearboxId
  time
  isClear
  hasSignificantErrors

2. Wheel
  id
  name
  isDefault

3. Pedals
  id
  name
  isDefault

4. Cockpit
  id
  name
  isDefault

5. Car
  id
  name
  image (optional)
  createdAt
  description (optional, make it field text input, maybe here will store some useful information about car, settings and so on)

6. Track
  id
  name
  image (optional)
  createdAt
  description (similar to cars)

7. Game
  id
  name
  createdAt
  image (optional) - just so end page will look better
