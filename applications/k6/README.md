### k6
O k6 é uma ferramenta de código aberto projetada para realizar testes de carga e estresse em aplicações, especialmente em APIs e microsserviços. Desenvolvida pela Grafana Labs, ela permite que desenvolvedores e engenheiros avaliem o desempenho e a confiabilidade de seus sistemas sob diferentes condições de carga.

[Site Oficial](https://k6.io/)

##### Instalar k6 
Para instalar k6, entre no site [Install k6](https://grafana.com/docs/k6/latest/)


##### Scripts
O k6 foi desenvolvido em Go, e ele tem uma linguagem javascript própria é uma variante do JavaScript, mais especificamente baseada no ECMAScript 2015 (ES6).

Foi feito 3 scripts para que possamos brincar e fazer alguns testes de stress.

##### Executar k6

```bash
$ k6 run -d 10s -u 10 post-get.ts --env PORT=8081
```

```bash
$ k6 run -d 10s -u 10 post-get.ts --env PORT=8080
```

```bash
$ k6 run -d 10s -u 10 post.ts --env PORT=8081
```
