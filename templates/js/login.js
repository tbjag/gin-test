// login.js

function handleLoginFormSubmission() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Example using Fetch API
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: username,
            password: password,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data && data.token) {
            // Redirect to the specified route
            window.location.href = data.redirect_url || '/default-route';
        } else {
            // Handle login failure
            console.error('Login failed');
        }
    })
    .catch(error => {
        console.error('Error during login:', error);
    });
}
