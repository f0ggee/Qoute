FROM golang:1.23



LABEL gmail=frederickscubi@gmail.com
ENV DATABASE_URL=postgresql://postgres.gilvehfffvzllxcicizz:zyjjef-sevpAp-5kitci@aws-0-eu-central-1.pooler.supabase.com:6543/postgres
WORKDIR /app
COPY . .

RUN go mod download &&  go mod tidy

ENTRYPOINT ["go", "run", "main.go"]


