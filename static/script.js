document.addEventListener('DOMContentLoaded', (event) => {
  console.debug('Script loaded')
  const map_name = document.getElementById('form-map-name')
  const map_title = document.getElementById('form-map-title')
  const map_version = document.getElementById('form-map-version')
  const grid_row = document.getElementById('form-grid-row')
  const grid_column = document.getElementById('form-grid-column')
  const terrain_type = document.getElementById('form-terrain-type')
  const cell_contents = document.getElementById('form-cell-contents')
  const response_output = document.getElementById('response')

  const save_button = document.getElementById('save-button')

  var mapCells = [[]]

  const load_map_list = async function() {
    try {
      const response = await fetch('/maps');
      const data = await response.json();

      const json_string = JSON.stringify(data, null, 2)

      response_output.innerHTML = json_string;

      data["maps"].forEach(optionText => {
        const newOption = new Option(optionText, optionText);
        map_name.add(newOption);
      });

      map_name.dispatchEvent(new Event('change'))

    } catch(e) {
      console.error(e)
    }
  }

  const load_map = async function(mapName) {
    try {
      // clear rows and columns in dropdown
      grid_row.innerHTML = ''
      grid_column.innerHTML = ''

      console.log(`map '${mapName}' selected`)

      const response = await fetch(`/maps/${mapName}`);
      const data = await response.json();

      const json_string = JSON.stringify(data, null, 2)

      response_output.innerHTML = json_string;

      map_title.value = data["map"]["title"]
      map_version.value = data["map"]["version"]

      mapCells = data["map"]["cells"]

      const row_count = mapCells.length
      const col_count = mapCells[0].length

      for(let i = 0; i < row_count; i++) {
        const newOption = new Option(i, i);
        grid_row.add(newOption);
      }
      for(let i = 0; i < col_count; i++) {
        const newOption = new Option(i, i);
        grid_column.add(newOption);
      }

      grid_column.dispatchEvent(new Event('change'))
    } catch(e) {
      console.error(e)
    }
  }

  const map_select_handler = async function(event) {
    const mapName = event.target.value;
    load_map(mapName)
  }

  const cell_select_handler = async function(event) {
    const row = grid_row.value
    const col = grid_column.value
    const cell = mapCells[row][col]

    console.log(`Cell (${row}, ${col}) selected - ${JSON.stringify(cell)}`)

    terrain_type.value = cell["type"]
    cell_contents.value = cell["contents"]
  }

  const save_handler = async function() {
    try {
      const post_data = {
        "type": terrain_type.value,
        "contents": cell_contents.value
      }
      const row = grid_row.value
      const col = grid_column.value
      const name = map_name.value

      const response = await fetch(`/maps/${name}/update/${row}/${col}`, {
        method: 'POST', // Specify the method
        headers: {
          'Content-Type': 'application/json' // Inform the server we are sending JSON
        },
        body: JSON.stringify(post_data) // Convert JS object to a JSON string
      });

      const data = await response.json();
      const json_string = JSON.stringify(data, null, 2)

      console.debug(`Response: ${json_string}`);
      response_output.innerHTML = json_string;

      load_map(name)
    } catch(e) {
      console.error(e)
    }
  }

  map_name.addEventListener('change', map_select_handler);
  grid_row.addEventListener('change', cell_select_handler);
  grid_column.addEventListener('change', cell_select_handler);

  save_button.addEventListener('click', function() {
    console.debug('Save button was clicked!');
    save_handler()
  });

  //load_handler();
  load_map_list();
})
