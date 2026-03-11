const API = "http://localhost:8080"
async function register() {
    const body = {
        name:     document.getElementById("reg-name").value.trim(),
        age:      parseInt(document.getElementById("reg-age").value.trim()),
        login:    document.getElementById("reg-login").value.trim(),
        password: document.getElementById("reg-password").value.trim(),
    }
    if (body.name || body.age || body.login || body.password)
    try {
        const res = await fetch(`${API}/register`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body)
        })

        if (!res.ok) {
            showMessage("reg-msg", "Такой логин уже существует", "error")
            return
        }
        
        const data = await res.json()
        showMessage("reg-msg", "Успешно! перенаправляем...", "success")
        setTimeout(() => {
            window.location.href = "/index.html"
        }, 1000)
    }
    catch (e) {
        showMessage('reg-msg', "Сервер недоступен", "error")
    }
}
function showMessage(id, text, type) {
    const el = document.getElementById(id)
    el.textContent = text
    el.className = `message ${type}`
}