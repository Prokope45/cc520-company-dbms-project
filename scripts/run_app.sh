# Run backend server
CURR_DIR="$(pwd)"
go run src/company/app/backend/server/main.go
# Run frontend server
cd src/company/app/frontend && npm run dev
cd $CURR_DIR