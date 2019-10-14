export HERO_ID=1
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json|jq -r '.[0].connections.relatives'
