fetch('http://localhost:8080/api/data')
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log(data); // Use the data here
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });

