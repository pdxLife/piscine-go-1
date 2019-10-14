export HERO_ID=52
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r ".[$HERO_ID].connections.relatives"
