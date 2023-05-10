import http from 'k6/http'

export let options = {
    vus: 8,
    duration: '3s'
}

export default function() {
    http.get("http://host.docker.internal:5000/health")
}