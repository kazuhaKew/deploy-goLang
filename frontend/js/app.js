document.addEventListener('DOMContentLoaded', function() {
    fetchMessage();

    document.getElementById('refresh-btn').addEventListener('click', fetchMessage);

    function fetchMessage() {
        const messageElement = document.getElementById('api-message');
        messageElement.textContent = 'Loading...';

        fetch('/api/hello')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                messageElement.textContent = data.message;
            })
            .catch(error => {
                console.error('Error fetching API data:', error);
                messageElement.textContent = 'Error loading message. Please try again.';
            });
    }
});