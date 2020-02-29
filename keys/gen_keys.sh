KEY_NAME="jwtRS256"
PUB_KEY=$(pwd)/$KEY_NAME.key.pub

if [[ ! -f "$PUB_KEY" ]]; then
  echo "generating key..."
  ssh-keygen -t rsa -b 4096 -f $KEY_NAME.key
fi

echo "setting key env variable (JWT_PUBLIC_KEY)..."
export JWT_PUBLIC_KEY=$(pwd)/$KEY_NAME.key.pub
