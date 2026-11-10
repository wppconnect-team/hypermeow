# Plano executavel: Cloud Go-only e HyperMeow WPPConnect

Inicio: 2026-11-10. Responsavel: joaosouz4dev. Atualizar os checkboxes somente com evidencia.
[x] significa concluido e verificado no escopo descrito; publicacao e testes reais tem itens separados.
Nao declarar 100% enquanto houver requisito ou validacao pendente. Preservar trabalho local existente.

## Fase 0 - Base e rastreabilidade
- [x] Comparar Go, Node e imagem Zuck v2.4.2 por revisoes/digests fixos.
- [x] Registrar autorizacao do usuario para publicar e fonte publica indicada: polymorfa/hypermeow.
- [x] Criar checkouts isolados Cloud/Manager/Go; preservar checkouts humanos.
- [x] Criar checkout isolado worker a partir da main 3ca6010.
- [x] Inventariar fonte publica HyperMeow, licencas, relacao com tulir e diferencas da imagem.
- [x] Confirmar se a imagem contem fonte privado completo; nao equiparar binario com fonte recuperado.

## Fase 1 - Cloud e pools somente Go
- [x] Cloud API e Manager rejeitam selecao diferente de Go; defaults e retorno coerentes.
- [x] Reservas do Manager filtram Node mesmo com configuracoes antigas.
- [x] Worker gera apenas Go, sem reposicao de Evolution/Wuzapi/Node por env legado.
- [x] Worker rejeita novas criacoes de outros providers antes de Docker/DB.
- [ ] Preservar operacao/limpeza de sessoes existentes sem reatribuir volumes ou credenciais.
- [ ] Testar defaults, env legado, pedido forjado, concorrencia e falhas de provisionamento.
- [ ] Identificar ambiente em uso, configurar reserva Go, publicar fontes testadas e verificar pool real.

## Fase 2 - Documentacao completa sem escolha de motor
- [x] Exemplos de criacao com JSON minimo, sem provider/runtime obrigatorios.
- [x] Guia publico: autenticacao/scopes, QR, status, texto, midia, lifecycle, limites e erros.
- [x] Corrigir rotas/payloads/eventos documentados conforme API real; distinguir recurso disponivel e pendente.
- [x] Publicar referencia operacional: pool, imagem, persistencia, backup, logs, rollout/rollback.
- [x] Atualizar traducoes PT/EN e verificar links/exemplos/build.

## Fase 3 - Projeto publico HyperMeow
- [x] Preparar repositorio wppconnect-team/hypermeow a partir de polymorfa/hypermeow, com historico/licencas.
- [x] Documentar origem, escopo SDK vs servidor Zuck e modulos/imports para uso no Go.
- [x] Rodar testes/vet/build e CI da base; registrar qualquer dependencia externa indisponivel.
- [x] Publicar repositorio publico e conferir visibilidade/SHA/licencas.
- [ ] Integrar atualizacoes de tulir preservando extensoes HyperMeow, com conflitos/testes resolvidos.
- [ ] Validar adaptacao do server-go ao SDK antes de trocar dependencia em producao.

## Fase 4 - Sincronizacao diaria
- [x] Agendar comparacao diaria com tulir/whatsmeow as 09:30 America/Sao_Paulo.
- [ ] Executar em checkout isolado, resolver principal no provedor e registrar SHAs de origem/destino.
- [x] Preservar patches e licencas; nunca usar reset/push force para igualar forks.
- [ ] Testar, incorporar e enviar atualizacoes compativeis; falhas ficam registradas sem promover.
- [x] Executar primeira comparacao agora e validar configuracao da tarefa agendada.

## Fase 5 - Integracao Cloud e confiabilidade
- [ ] Corrigir webhook por sessao na ativacao do pool, inclusive webhookUrl/webhook, restart e isolamento.
- [x] Implementar HMAC opcional nos webhooks globais do Go, com testes e documentacao.
- [ ] Idempotencia persistida de envio e recuperacao de resultado ambiguo.
- [ ] Falhas de entrega: DLQ, filtros, replay auditado, metricas e alarmes de backlog.
- [ ] Testes reais de QR/phone, reconexao, envio/recebimento/midia/ACK/logout e dois tenants.

## Fase 6 - Paridade funcional Zuck
- [ ] Converter inventario de 303 operacoes em matriz executavel de contratos e criterios por jornada.
- [ ] SDK: incorporar recursos publicos HyperMeow e identificar extensoes privadas nao presentes.
- [ ] Adaptadores de contatos, grupos/comunidades, newsletters, status e midia.
- [ ] Integracoes Chatwoot/n8n, SQS/Redis Streams e armazenamento automatico de midia.
- [ ] Comercio: catalogos/produtos/pedidos/pagamentos com contratos e idempotencia.
- [ ] Interativos modernos: botoes/listas/carrosseis/flows com teste Android/iOS/Web.
- [ ] Chamadas: sinalizacao/audio/video/grupos, autenticacao de streams e limites de recursos.
- [ ] Testes de protocolo para cada recurso; ausencia de fonte ou teste nao conta como paridade.

## Fase 7 - Aceite e encerramento
- [ ] CI, race, integracoes PostgreSQL/brokers/S3, build e smoke multiarch aprovados.
- [ ] Testes de carga/falha/restauracao com resultados registrados e metas aceitas.
- [ ] Repositorios e ambientes publicados a partir das fontes verificadas; health/logs/versao confirmados.
- [ ] Checklist integral concluido com links de evidencia. Somente entao declarar 100%.

## Evidencias
- Analise: server-go/docs/research/server-go-parity-2026-11-10.md.
- Primeira implementacao: IMPLEMENTATION.md.
- Base publica/privada: public-hypermeow-foundation.md (concluido).
- Este plano e o registro principal; complementar com recibos de testes/publicacao sem incluir segredos.

## Registro de validacao - 2026-11-10
- HyperMeow publico: https://github.com/wppconnect-team/hypermeow, main bf5bab849ea6f666f490e32487cec0a7c84dbeb5. Historico e licencas preservados.
- CI HyperMeow aprovado: https://github.com/wppconnect-team/hypermeow/actions/runs/42860662627 (Go antigo/atual, testes, build, WASM, MEX e pre-commit).
- Automacao ativa: sincronizar-hypermeow-wppconnect-diariamente, diariamente 09:30 America/Sao_Paulo. Configuracao conferida pela ferramenta do aplicativo. As regras da fase 4 estao no prompt; a primeira integracao ainda tem 15 conflitos a resolver.
- Auditoria das 10 camadas AMD64: imagem nao contem fonte privado completo; contem binario, UI e Swagger. O projeto publico nao e uma recuperacao desse fonte.
- Cloud: 43 testes locais e build Next.js aprovados. PR https://github.com/joaosouz4dev/wppconnect-cloud/pull/37.
- Worker: 13 testes, 3 smokes e build TypeScript aprovados. PR https://github.com/joaosouz4dev/wppconnect-cloud-manager-worker/pull/6.
- Manager: 14 testes unitarios aprovados; integracoes locais nao executaram por falta de PostgreSQL (ECONNREFUSED). CI com banco isolado pendente. PR https://github.com/joaosouz4dev/wppconnect-cloud-manager/pull/5.
- Guia operacional esta versionado nos PRs; publicacao no ambiente Cloud continua pendente na fase 1.
