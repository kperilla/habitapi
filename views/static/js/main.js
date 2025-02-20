var apiBaseUrl = ''
makePostForm('userForm', '/api/v1/users/')
makePostForm('habitGroupForm', '/api/v1/habit_groups/')
// makePostForm('userForm', '/api/v1/users/')
// makePostForm('userForm', '/api/v1/users/')

function makePostForm(formId, resourcePath, baseUrl=apiBaseUrl) {
    document.getElementById(formId).addEventListener('submit', async function (event) {
        event.preventDefault(); // Prevents the default form submission

        // Convert form data to a JSON object
        const formData = new FormData(this);
        const jsonObject = Object.fromEntries(formData.entries());

        try {
            requestUrl = baseUrl + resourcePath
            const response = await fetch(requestUrl, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(jsonObject),
            });

            if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);
            const result = await response.json();
            console.log('Success:', result);
        } catch (error) {
            console.error('Error:', error.message);
        }
    });
}
