async function getMapList() {
    try {
        const response = await fetch('/maps');
        const data = await response.json();

        return data["maps"]
    } catch(e) {
        console.error(e)
        return null
    }
}

async function loadMap(mapName) {
    try {
        const response = await fetch(`/maps/${mapName}`);
        const data = await response.json();

        return data["map"]
    } catch(e) {
        console.error(e)
        return null
    }
}

async function saveHex(mapName, row, col, terrainType, contents) {
    try {
        const post_data = {
            "type": terrainType,
            "contents": contents
        }

        const response = await fetch(`/maps/${mapName}/update/${row}/${col}`, {
            method: 'POST', // Specify the method
            headers: {
                'Content-Type': 'application/json' // Inform the server we are sending JSON
            },
            body: JSON.stringify(post_data) // Convert JS object to a JSON string
        });

        const data = await response.json();

        return data
    } catch(e) {
        console.error(e)
        return null
    }
}
