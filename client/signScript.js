window.addEventListener("load", () => {
  function sendData() {
    const XHR = new XMLHttpRequest();

    // Bind the FormData object and the form element
    const emailObjectValue = document.getElementById("inputEmail").value
    const passObjectValue = document.getElementById("inputPassword").value

    const data = {
      email: emailObjectValue,
      password: passObjectValue
    }

    // Define what happens on successful data submission
    XHR.addEventListener("load", (event) => {
      const element = document.getElementById("f").parentNode
      element.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="10rem" height="10rem" fill="green" class="bi bi-check-circle-fill" viewBox="0 0 16 16">
      <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zm-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z" />
    </svg > <p class="fs-1 fw-bald text-center text-uppercase">successful</p>`
      element.classList.add("text-center")
      setTimeout(() => {
        window.location.replace("http://localhost:8080");
      }, "1000")
    });

    // Define what happens in case of error
    XHR.addEventListener("error", (event) => {
      alert('Oops! Something went wrong.');
    });

    // Set up our request
    XHR.open("POST", "http://localhost:8080/api/user");

    // The data sent is what the user provided in the form
    XHR.send(JSON.stringify(data));

  }

  // Get the form element
  const form = document.getElementById("f");
  // Add 'submit' event handler
  form.addEventListener("submit", (event) => {
    event.preventDefault();

    sendData();
  });
});