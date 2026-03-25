document.addEventListener('DOMContentLoaded', (event) => {
  console.debug('Script loaded')
  const save_button = document.getElementById('save-button')
  const map_name = document.getElementById('form-map-name')
  const map_title = document.getElementById('form-map-title')
  const map_version = document.getElementById('form-map-version')
  const grid_row = document.getElementById('form-grid-row')
  const grid_column = document.getElementById('form-grid-column')
  const grid_value = document.getElementById('form-grid-value')
  const response_output = document.getElementById('response')

  var mapCells = [[]]

  const load_map_list = async function() {
    try {
      const response = await fetch('/maps');
      const data = await response.json();

      const json_string = JSON.stringify(data, null, 2)

      console.debug(`Response: ${json_string}`);
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

  const map_select_handler = async function(event) {
    try {
      const mapName = event.target.value;

      // clear rows and columns in dropdown
      grid_row.innerHTML = ''
      grid_column.innerHTML = ''

      console.log(`map '${mapName}' selected`)

      const response = await fetch(`/maps/${mapName}`);
      const data = await response.json();

      const json_string = JSON.stringify(data, null, 2)

      console.debug(`Response: ${json_string}`);
      response_output.innerHTML = json_string;

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

     // grid_value.value = data['grid-cell']['contents']
    } catch(e) {
      console.error(e)
    }
  }

  const save_handler = async function() {
    try {
      const post_data = {
        'id': grid_id.value,
        'contents': grid_value.value
      }
      const response = await fetch('/api/save', {
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
    } catch(e) {
      console.error(e)
    }
  }

  map_name.addEventListener('change', map_select_handler);

  save_button.addEventListener('click', function() {
    console.debug('Save button was clicked!');
    save_handler()
  });

  //load_handler();
  load_map_list();
})
