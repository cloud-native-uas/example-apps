const express = require('express');
const axios = require('axios');
const app = express();

app.get('/forks', async (req, res) => {
    try {
        const response = await axios.get('https://api.github.com/repos/kubernetes/kubernetes');
        const forks = response.data.forks_count;
        res.send(`The Kubernetes repository has ${forks} forks.`);
    } catch (error) {
        console.error(error);
        res.send('Failed to fetch the number of forks');
    }
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});