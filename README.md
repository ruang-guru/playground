# Ruangguru Playground

[![GitHub Super-Linter](https://github.com/ruang-guru/playground/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)

![banner](banner.png)

---

## To generate repo without answer

- `go run cli/main.go answer remove -r backend` or `go run cli/main.go answer remove -r frontend`
- `rsync -vhra . ../playground --include='**.gitignore' --exclude='/.git' --filter=':- .gitignore' --delete-after`
- `cd ../playground`
- `git add .`
- `git commit -nm "(sync)"`
- `git push`
