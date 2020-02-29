KEY_NAME="jwtRS256"
echo "KEY_NAME=$KEY_NAME"
PWD=$(pwd)
echo "PWD=$PWD"
PUB_KEY="$PWD/$KEY_NAME.key.pub"
echo "PUB_KEY=$PUB_KEY"

if [[ ! -f "$PUB_KEY" ]]; then
  echo "generating key..."
  ssh-keygen -t rsa -b 4096 -f $KEY_NAME.key
fi

echo "setting key env variable (JWT_PUBLIC_KEY)..."
JWT_PUBLIC_KEY="$PWD/$KEY_NAME.key.pub"
echo "JWT_PUBLIC_KEY=$JWT_PUBLIC_KEY"
export JWT_PUBLIC_KEY
