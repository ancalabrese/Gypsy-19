import Axios from 'axios'
import Oktokit, { Octokit } from '@octokit/rest'


const OctokitOptions = {
    baseUrl: 'https://api.github.com',
    userAgent: '`${process.env.REACT_APP_NAME} ${process.env.REACT_APP_VERSION}`'
}

const GithubClient = new Octokit(OctokitOptions)

GithubClient.projects.get()

export default GithubClient;