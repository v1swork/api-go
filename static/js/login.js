const API = 'http://localhost:8080'

async function login() {
    const body = {
        login: document.getElementById("login-login").value,
        password: document.getElementById("login-password").value
    }

    try {
        // const res = await fetch('http://localhost:8080/users')
        // const data = await res.json()
        const res = await fetch(`${API}/login`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body)
            // {"login":"Ravil", "password": "123"}
        })
        const data = await res.json()

        if (res.ok) {
            localStorage.setItem("token", data.token)
            showMessage("login-msg", "Успешный вход. Перенаправляем...", "success")
            setTimeout(() => {
                window.location.href = "/users.html"
            }, 1000)
        } else {
            showMessage("login-msg", "Неверный логин или пароль", "error")
        }
    } catch (e) {
        showMessage("login-msg", "Сервер недоступен", "error")
    }
}

function showMessage(id, text, type) {
    const el = document.getElementById(id)
    el.textContent = text
    el.className = `message ${type}`
}