import http from 'k6/http'

export let options = {
    stages:[
        {target: 1, duration: '2s'},
        {target: 15, duration: '10s'},
        {target: 3, duration: '2s'},
    ]
}

export default function() {
    http.get("http://host.docker.internal:5000/v1/product/getProducts")
}