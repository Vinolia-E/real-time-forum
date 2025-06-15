const form = document.getElementById("registerForm");
form.addEventListener("submit", function(e) {
   e.preventDefault();
   fetch("http://localhost:8080/register", {
    method: "POST",
    body: new FormData(form),
   })
   .then(response => response.json())
   .then(data => {
    console.log("Success:", data.message); 
   })
   .catch((error) => {
    console.error("Error:", error);
    });
}); 