const API = "http://localhost:8080"
const token = this.localStorage.getItem("token")

window.onload = function() {
    if (!token) {
        this.alert("Авторизуйтесь сначала")
        window.location.href = "index.html"
    } else {
        getUsers()
    }
}
async function getUsers() {
    try {
        const res = await fetch(`${API}/users`, {
            method: "GET",
            headers: {"Authorization": `Bearer ${token}`}
        })
        if (res.status === 401) {
            localStorage.removeItem("token")
            window.location.href = "index.html"
            return
        }

        if (!res.ok){
            showMessage("users-msg", "Ошибка загрузки", "error")
            return
        }

        const data = await res.json()
        const list = document.getElementById("users-list")
        list.innerHTML = ""

        data.forEach(user => {
            list.innerHTML += `
            <div class = "user-card">
                <b>ID:</b> ${user.id} |
                <b>Имя:</b> ${user.name} |
                <b>Возраст:</b> ${user.age} |
                <b>Логин:</b> ${user.login}
            </div>
            `    
        })
    } catch (e) {
        showMessage('users-msg', "Сервер недоступен", "error")
    }
}

function logout() {
    localStorage.removeItem("token")
    window.location.href = "index.html"
}

function showMessage(id, text, type) {
    const el = document.getElementById(id)
    el.textContent = text
    el.className = `message ${type}`
}