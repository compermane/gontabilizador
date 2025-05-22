# Gontabilizador
Trata-se de um contabilizador de presenças em ensaios para ritmistas. Tanto a comunicação com o banco de dados como o servimento e carregamento de arquivos estáticos (html, css) são realizados por **Go**. O banco de dados foi construído com **MySQL**.

Além disso, também há um script em **Python** que escuta por atualizações no banco de dados, escrevendo no terminal atualizações de inserção na tabela de presença.

## Uso
Primeiramente, faça o clone do repositório
```bash
git clone https://github.com/compermane/gontabilizador
```

Feito o clone, rode o seguinte comando do Docker
```bash
docker-compose up --build
```
Esse comando vai construir imagens definidas nos respectivos Dockerfile e orquestrar os contêineres para que se comuniquem corretamente. O processo leva alguns minutos. Ao ser finalizado, o projeto fica disponível em internamente no docker por `gontabilizador-app:3000`.

## Desinstalando
As imagens geradas ocupam espaço considerável em disco (>2GB). Portanto, após o uso, é recomendado que este projeto seja desinstalado. Para tanto, execute
```bash
docker-compose down -v
```
Isso "derruba" as imagens geradas e apaga seus volumes. Após isso, remova todos as imagens geradas com
```bash
docker rmi -f <nome-ou-id-da-imagem>
```

### Autor
Eugênio Akinori Kisi Nishimiya RA: 811598
