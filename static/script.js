document.addEventListener('DOMContentLoaded', (event) => {
  console.debug('Script loaded')
  const load_button = document.getElementById('load-button')
  const save_button = document.getElementById('save-button')
  const grid_id = document.getElementById('form-grid-id')
  const grid_value = document.getElementById('form-grid-value')
  const response_output = document.getElementById('response')

  const load_handler = async function() {
    try {
      const response = await fetch('/api/load');
      const data = await response.json();

      const json_string = JSON.stringify(data, null, 2)

      console.debug(`Response: ${json_string}`);
      response_output.innerHTML = json_string;

      grid_id.value = data['grid-cell']['id']
      grid_value.value = data['grid-cell']['contents']
    } catch(e) {
      console.error(e)
    }
  }

  save_button.addEventListener('click', function() {
    console.debug('Save button was clicked!');
  });

  load_button.addEventListener('click', function() {
    console.debug('Load button was clicked!');
    load_handler()
  });

  load_handler();
})
