export HERO_ID=0
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json|jq -r '.[52].connections.relatives'
