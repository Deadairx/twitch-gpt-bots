import 'dotenv/config';
import tmi from 'tmi.js';

const options: tmi.Options = {
    identity: {
        username: process.env.BOT_USERNAME || '',
        // TODO: get oauth token from https://twitchapps.com/tmi/
        password: process.env.OAUTH_TOKEN || '',
    },
    channels: [process.env.CHANNEL_NAME || ''],
};

const client: tmi.Client = new tmi.Client(options);

client.connect();

client.on('message', (channel, userstate, message, self) => {
    if (self) return;

    console.log(`Received message from ${userstate.username}: ${message}`);

    // TODO: add commands

    // TODO: Filter for messages that mention the bot

    if (message.toLowerCase() === '!hello') {
        client.say(channel, `@${userstate.username}, heya!`);
    }
});

client.on('disconnected', (reason) => {
    console.log(`Disconnected: ${reason}`);
    process.exit(1);
});

