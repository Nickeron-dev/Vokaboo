import discord
from keep_alive import keep_alive

#load_dotenv()
TOKEN = 'OTc5Mjk2NzAwNTMzNDQ4NzI0.GKarIj.t8sWwQ7b9aBuF_wh77eKGHqvQOfWbMoqYWIszQ'

client = discord.Client()

@client.event
async def on_ready():
    print(f'{client.user} has connected to Discord!')
    channel = client.get_channel(979297858815668227)
    await channel.send('hello')

@client.event
async def on_message(message):
    if message.content.startswith('$greet'):
        channel = message.channel
        await channel.send('Say hello!' + message.content)

keep_alive()
client.run(TOKEN)

