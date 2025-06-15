fetch("http://localhost:8080/register", {
    method: "POST",
    body: new FormData(form),
})
.then(response => response.json())
.then(data => {
    console.log("Success:", data.message); // or display it in the UI
})
.catch(error => {
    console.error("Error:", error);
});
