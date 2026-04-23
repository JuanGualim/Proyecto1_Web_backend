async function exportCSV() {
    const response = await fetch("http://localhost:8080/series")
    const data = await response.json()

    let csv = "ID,Name,Current Episode,Total Episodes\n"

    data.forEach(series => {
        csv += `${series.id},${series.name},${series.current_episode},${series.total_episodes}\n`
    })

    const blob = new Blob([csv], { type: "text/csv" })
    const url = window.URL.createObjectURL(blob)

    const a = document.createElement("a")
    a.href = url
    a.download = "series.csv"
    a.click()

    window.URL.revokeObjectURL(url)
}