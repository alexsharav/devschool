async function checktoken() {
  try {
    let response = await fetch("http://localhost:8080/accesschecker", {
      credentials: "include",
    });

    if (response.ok) {
      const data = await response.json();
      return data.authenticated === true;
    }

    const refreshResponse = await fetch("http://localhost:8080/refresh", {
      credentials: "include",
    });

    if (refreshResponse.ok) {
      const secondResponse = await fetch(
        "http://localhost:8080/accesschecker",
        {
          credentials: "include",
        }
      );

      if (secondResponse.ok) {
        const data = await secondResponse.json();
        return data.authenticated === true;
      }
    }

    return false;
  } catch (error) {
    console.error("Token check error:", error);
    return false;
  }
}

export { checktoken };
