import Axios from 'axios'

const GithubClient = Axios.create({
    baseURL: 'https://api.github.com/',
    headers = {'Accept': 'application/vnd.github.v3+json'}
})