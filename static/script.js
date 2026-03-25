document.addEventListener('DOMContentLoaded', (event) => {
  console.debug('Script loaded')
  const load_button = document.getElementById('load-button')
  const save_button = document.getElementById('save-button')
  const map_name = document.getElementById('form-map-name`')
  const grid_row = document.getElementById('form-grid-row')
  const grid_column = document.getElementById('form-grid-column')
  const grid_value = document.getElementById('form-grid-value')
  const response_output = document.getElementById('response')

  const load_handler = async function() {
    try {
      const response = await fetch('/api/load');
      const data = await response.json();

      const json_string = JSON.stringify(data, null, 2)

      console.debug(`Response: ${json_string}`);
      response_output.innerHTML = json_string;

      grid_value.value = data['grid-cell']['contents']
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

  save_button.addEventListener('click', function() {
    console.debug('Save button was clicked!');
    save_handler()
  });

  load_button.addEventListener('click', function() {
    console.debug('Load button was clicked!');
    load_handler()
  });

  load_handler();
})
