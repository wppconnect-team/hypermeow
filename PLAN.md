# Plano executavel: Cloud Go-only e HyperMeow WPPConnect

Inicio: 2026-11-10. Responsavel: joaosouz4dev. Atualizar os checkboxes somente com evidencia.
[x] significa concluido e verificado no escopo descrito; publicacao e testes reais tem itens separados.
Nao declarar 100% enquanto houver requisito ou validacao pendente. Preservar trabalho local existente.
Os registros abaixo sao cronologicos: pendencias de um marco antigo podem estar encerradas nos marcos seguintes. O checklist e o ultimo registro indicam o estado atual.

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
- [x] Preservar operacao/limpeza de sessoes existentes sem reatribuir volumes ou credenciais.
- [x] Testar defaults, env legado, pedido forjado, concorrencia e falhas de provisionamento.
- [x] Identificar ambiente em uso, configurar reserva Go, publicar fontes testadas e verificar pool real.

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
- [x] Integrar atualizacoes de tulir preservando extensoes HyperMeow, com conflitos/testes resolvidos.
- [x] Validar adaptacao do server-go ao SDK antes de trocar dependencia em producao.

## Fase 4 - Sincronizacao diaria
- [x] Agendar comparacao diaria com tulir/whatsmeow as 09:30 America/Sao_Paulo.
- [x] Executar em checkout isolado, resolver principal no provedor e registrar SHAs de origem/destino.
- [x] Preservar patches e licencas; nunca usar reset/push force para igualar forks.
- [x] Testar, incorporar e enviar atualizacoes compativeis; falhas ficam registradas sem promover.
- [x] Executar primeira comparacao agora e validar configuracao da tarefa agendada.

## Fase 5 - Integracao Cloud e confiabilidade
- [x] Corrigir webhook por sessao na ativacao do pool, inclusive webhookUrl/webhook, restart e isolamento.
- [x] Implementar HMAC opcional nos webhooks globais do Go, com testes e documentacao.
- [x] Idempotencia persistida de envio, recuperacao de resposta e bloqueio de reenvio ambiguo; incerteza do protocolo continua explicita.
- [x] Falhas de entrega: consulta por sessao/destino/estado, replay seletivo auditado e metricas de backlog/capacidade, com testes de falha e recuperacao.
- [ ] ADIADO pelo usuario em 10/11/2026: nenhum monitoramento ou destino de notificacao por ora. Nao instalar nem disparar alertas operacionais; as regras seguem como exemplos e nao ha entrega validada.
- [x] Conta de teste real: QR, reconexao, texto/arquivo, ACK de entrega/leitura, resposta recebida, idempotencia e download de midia depois de restart. Outro workspace recebe 404.
- [ ] Logout e pareamento por codigo de telefone separados do QR.
- [ ] ADIADO pelo usuario em 10/11/2026: jornadas com segunda conta pareada e validacao Android/iPhone/Web. Nao bloqueiam a implementacao automatizada; nao contam como aprovadas.

## Fase 6 - Paridade funcional Zuck
- [x] Converter inventario de 303 operacoes em matriz executavel de contratos e criterios por jornada.
- [x] SDK: incorporar recursos publicos HyperMeow e identificar extensoes privadas nao presentes.
- [ ] Adaptadores de contatos, grupos/comunidades, newsletters, status e midia.
- [x] J10: corrigir espera/cancelamento de ACK de visualizacoes no SDK antes de expor a operacao.
- [x] J10 no server-go (contratos locais com cliente de teste): convites, reacoes, visualizacoes e atualizacoes de mensagens com contratos estritos, isolamento e idempotencia nas mutacoes.
- [x] J10: protocolo de ACK no SDK, fixtures HTTP/SQLite no Go, OpenAPI e guia nativo; CI aprovado e fontes incorporadas.
- [x] J10 na Cloud: escopos separados para leitura, escrita e administracao, resultados isolados, documentacao PT/EN e testes de navegador/build/CI.
- [x] J10: publicar a imagem no worker e renovar/verificar as tres reservas em Go; smoke pelo Manager sem telefone pareado.
- [x] J10: Cloud 5b9f3e8 publicada; /docs em portugues com o contrato de canais, quatro rotas protegidas responderam 401 e rota desconhecida 404.
- [ ] J10: aceite de owner/subscriber em duas contas e nos clientes Android/iPhone/Web adiado pelo usuario.
- [x] Redis Streams: rejeicao/recuperacao, trim/exact maxlen e identidade do evento com Redis real no CI; incorporado em [Go #22](https://github.com/wppconnect-team/wppconnect-server-go/pull/22).
- [x] SQS: destino opcional, rejeicao de fila inexistente, recuperacao e deduplicacao FIFO com ElasticMQ; incorporado em [Go #24](https://github.com/wppconnect-team/wppconnect-server-go/pull/24). Sem teste de credenciais/IAM/KMS AWS reais.
- [x] n8n: importar workflow e credencial temporarios, publicar, recusar segredo errado e processar evento autenticado em n8n 2.47.3 isolado ([CI](https://github.com/wppconnect-team/wppconnect-server-go/actions/runs/42902244064/job/130130987174)).
- [ ] Integracao Chatwoot e armazenamento automatico de midia. CRUD/segredos de brokers por sessao (J15) continuam separados dos adaptadores globais opcionais.
- [x] J11 no server-go (contrato local): catalogo proprio, produto e colecoes, limite e cursor, escopo da conta e preco sem conversao.
- [x] J11: fontes de leitura incorporadas em [Go #23](https://github.com/wppconnect-team/wppconnect-server-go/pull/23), CI com race e imagens.
- [x] J11: leitura publicada na Cloud, protecao das tres rotas, guia EN/PT e tres reservas Go saudaveis; [recibo](docs/brokers-catalog-acceptance.json).
- [ ] J11: aceite de catalogo em conta Business com produtos/colecoes reais.
- [ ] Comercio: catalogos/produtos/pedidos/pagamentos com contratos e idempotencia.
- [x] J05/Z126: contrato HTTP experimental de carrossel PNG/JPEG com respostas rapidas; validacao antes de upload, isolamento, deduplicacao e payload protobuf testados; CI 42901305610 aprovado e main 93304d79bd09af36386035170e78886e6b151e56.
- [ ] Interativos modernos: botoes/listas/carrosseis/flows com teste Android/iOS/Web.
- [ ] Chamadas: sinalizacao/audio/video/grupos, autenticacao de streams e limites de recursos.
- [ ] Testes de protocolo para cada recurso; ausencia de fonte ou teste nao conta como paridade.

## Fase 7 - Aceite e encerramento
- [x] CI, race, integracoes PostgreSQL/brokers/S3, build e smoke multiarch aprovados para as mudancas incorporadas ate Go e01698b e Cloud 35ac0c4.
- [x] Restauracao offline do backup consistente da conta de teste: integridade dos cinco bancos SQLite, sem abrir uma segunda conexao WhatsApp.
- [x] Ensaio isolado de recuperacao idempotente com 1 e 4 clientes: 72 leituras autenticadas identicas, quatro registros preservados depois de restart e quatro consultas entre sessoes negadas.
- [ ] Carga de protocolo, reconexao, midia e backlog com metas de capacidade aceitas; o microbenchmark HTTP nao cobre esses perfis.
- [x] Mudancas implementadas publicadas a partir das fontes verificadas; health/logs/versao confirmados em 10/11/2026.
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

## Publicacao Cloud Go-only - 2026-11-10
- Cloud main 70c00e087ccc12d0235fb0ee0ab58dbb7f16dec7, deployment dpl_QKxcN2iGNJZ1LF2Xq6raKu9JWfVg. Dominio confirmado; /docs HTTP 200, criacao anonima HTTP 401. Evidencia cloud-release.json.
- Manager main 3ade44f97868c008c30a1e0894f9636cfa851c51; worker main f87a9a9b9d18579f5f1aa8521a125c75a8c3e9cb. CI aprovado, inclusive PostgreSQL, concorrencia e rollback. Imagens e versoes anteriores em go-only-services-release.json.
- Canary: rejeicao de criacao nao-Go no worker/Manager, autenticacao, bancos e proxy da sessao existente aprovados. Containers temporarios removidos; evidencia go-only-service-canary.json.
- Producao: WPP_RUNTIME_POOLS=go:go=3, tres reservas Go, pedido Node HTTP 422, sessao antiga CONNECTED. Outros containers do host preservados.
- Smoke real com workspace isolado: criacao minima HTTP 200/Go, QR PNG valido, outro tenant HTTP 404, cota HTTP 409, exclusao de teste concluida e reposicao automatica para tres reservas Go. Nao houve envio de mensagens nem pareamento de conta de cliente.
- Integracao HyperMeow com tulir em PR https://github.com/wppconnect-team/hypermeow/pull/1; 15 conflitos resolvidos e testes locais aprovados; CI/race pendentes. Ainda nao promovida a main.

## SDK e contratos - 2026-11-10
- HyperMeow sincronizado com tulir c88388f: PR #1 incorporado, merge aae5d9a62408c202fb4dd2f91394153c054a110b. CI Go 1.26/1.27, race, WASM, MEX e pre-commit aprovados nas runs 42866085986 e 42866095663.
- HMAC server-go: PR #15 incorporado, merge 36e41bc; CI 42865161731 aprovado com race, brokers/S3 e smoke amd64/arm64. Ainda nao altera a imagem 0.1.11 da pool.
- Adapter HyperMeow server-go: PR #16 incorporado, merge dc8d124. CI 42866854308 aprovado com race, integracoes e imagens das duas arquiteturas. Modulo remoto fixado em 0987e7d211c8.
- Teste SQLite reproduziu duplicate column na migracao da base anterior v16. Corrigido no SDK e validado com preservacao de chaves Noise/identidade/ADV, segredo WASA, indice faltante e upgrade repetido.
- Matriz estruturada com 303 operacoes e 20 jornadas: parity-contract-matrix.md / .json. Versao publica compacta lossless 482233 bytes, validada por expansao e comparacao integral. Todas as operacoes seguem functionalParityValidated:false no baseline pesquisado; nao sao testes de mensagens reais.
- Webhook por sessao: implementacao em branches server-cloud-webhooks e worker-webhooks, com receiver dedicado, segredo e restricao de egress HTTPS publico; rollout e verificacao ponta a ponta pendentes.

## Webhooks publicados e verificados - 2026-11-10
- Server-go PR #17 e worker PR #7 incorporados apos CI aprovado. Go main 0878a10e51beae17fa199e2abd9b76ddf0f59ac1; worker main 95aa5d279778e84707b5bfde886cdd455bee48fe.
- Imagens construidas de arquivos git de checkouts limpos e hashes conferidos. Go sha256:e249ddc5d0f6a06d65cce3e63e178cc25b68cbf597b0bf41871b869430a21e8a; worker sha256:031ccfa74a61b0e176210d0cbefafdf4e5f7b21dca363b4ee392eea953605fef.
- Canary real: dois receptores HTTPS com chaves distintas, assinatura validada sobre corpo original, primeira entrega 503 seguida de retry com ID estavel, zero assinaturas invalidas, nenhum evento cruzado, restart preservado. Producao e pool preservadas. Evidencia webhooks-canary.json.
- Rollout: GO_SESSION_WEBHOOKS_ENABLED=true e imagem nova para futuras criacoes. Tres reservas anteriores e sessao Node CONNECTED preservadas. Backup privado registrado em webhooks-release.json.
- Manager real: criacao com webhook em 5,91 segundos, QR PNG, webhook assinado, outro workspace 404, cota 409, exclusao concluida. Nenhum telefone pareado. Evidencia webhooks-live-smoke.json.
- Receptor HTTPS temporario, rota e segredos de teste removidos. Evidencia webhooks-cleanup.json.
- Cloud PR #38 incorporado. Documentacao PT/EN publicada em https://cloud.wppconnect.io/docs; HTTP 200 e exemplo webhookSecret confirmado. Deployment dpl_Bdf6d87gR9EaGcUjEtmnizxoQ73j, fonte e895a7d. 43 testes, build e CI aprovados.
- HyperMeow PR #2 incorporado: matriz publica completa e testes herdados restaurados, CI Go antigo/atual aprovado.
- Proxima etapa em andamento: ledger persistido de comandos de envio, com chave por sessao, recuperacao de resposta e bloqueio de reenvio ambiguo. Branch server-idempotency; testes focados aprovados, suite completa/CI/rollout ainda pendentes.

## Criterios de fechamento das lacunas restantes
- Idempotencia: concorrencia, reinicio, falha de persistencia, timeout parcial, isolamento, rejeicao em imagens antigas e transporte do header por todos os proxies; depois smoke no artefato publicado.
- DLQ: listar entregas falhas por sessao/destino, motivo sanitizado, replay seletivo sem repetir destinos concluidos, auditoria, metrica de idade/ocupacao e ensaio de indisponibilidade/recuperacao.
- Paridade: implementar por jornada da matriz, com fixtures de protocolo e contrato; comparar resultado observavel e nao somente nome de rota. Marcar cada operacao somente depois de validar comportamento.
- WhatsApp real: destinatario autorizado recebido; depende agora de conta de teste pareada. Nunca usar sessoes de clientes para envios de teste.
- Atualizacao de aceite: destinatario autorizado recebido em 10/11/2026 e mantido em registro privado, sem publicar o telefone. Instancia Go isolada criada na imagem validada; QR apresentado ao usuario. Pareamento ainda pendente, nenhuma mensagem real enviada.
- Fonte privada: todas as camadas da imagem foram inspecionadas e nao contêm o fonte completo. As extensoes publicas foram preservadas; recursos exclusivamente binarios precisam de implementacao propria ou fornecimento do fonte. Nao declarar clonagem integral.

## Confiabilidade incorporada - 2026-11-10
- Server-go PR #18: ledger persistido de envio com chave por sessao, fingerprint do corpo/rota/Content-Type, resposta recuperavel, reserva antes do efeito de protocolo, bloqueio de reenvio em resultado incerto e limites de armazenamento. CI 42873726008 aprovado, incluindo race e imagens multiarch. Merge c29f4ac.
- Server-go PR #19: filtros de entregas por destino/estado, motivo sanitizado, retry somente de destino falho, auditoria transacional, migracao da base anterior e metricas de idade/capacidade. CI 42875435897 aprovado. Main 514347cd5e7dd69348a1ef1d63e51f31ae29cf14.
- Cloud PR #39: preservacao de Idempotency-Key, consulta send-result e verificacao previa do ledger para recusar imagens antigas sem enviar. 48 testes, build e CI aprovados. Main 35ac0c4080dd375a3efa1625567e55d93c856958; deployment dpl_ATXHSjh4bJDtWKBKoQC5CrvT4M6R publicado e /docs conferido HTTP 200.
- Imagem Go sha256:e2c292c3259007d8d74c72422b847510ab2c75d43692e5dfb265345523dcb2ec, construida da main limpa e arquivo-fonte SHA256 c0c105975fab5eb167cedfdf2fac1b124480d2c775e7179846d283e65a8b86c3.
- Canary real da imagem: falhas HTTP repetidas, estado failed/retry_exhausted, recuperacao seletiva com assinatura valida, auditoria, bloqueio entre sessoes, metricas, resultado idempotente, conflito de payload e persistencia apos restart. Nenhuma mensagem enviada e nenhum telefone pareado. Fixtures e rede temporaria removidas; producao preservada.
- Primeiro ensaio abortou durante conexao inicial reiniciada no boot. O script foi corrigido para aguardar esse estado transitorio; ensaio completo seguinte aprovado. Nao foi uma falha funcional ocultada.
- Regras Prometheus foram documentadas como exemplos, sem declarar monitoramento instalado. Essa ativacao permanece pendente e separada da implementacao de metricas.
- Rollout e renovacao das tres reservas concluidos via alocacoes de teste proprias, sem apagar reservas diretamente. Todas usam a imagem e2c292c3. Smoke pelo Manager confirmou transporte da chave, recuperacao de resposta e conflito de payload; alocacoes de teste excluidas. Recibo reliability-live-pool.json.
- Auditoria final 18:46 UTC: Manager/worker em execucao, tres Go healthy, nenhum OOM, panic/fatal/falha de persistencia nos dez minutos consultados, sessao antiga CONNECTED e containers temporarios ausentes. Recibo reliability-final-audit.json.
- Esses smokes usaram contas nao pareadas e payloads invalidos para nao enviar mensagens reais. O protocolo de envio foi exercitado com clientes simulados nos testes; recebimento real em telefone permanece pendente.

## Sequencia da fase de paridade funcional
1. Contatos, grupos/comunidades e newsletters: partir dos adaptadores ja existentes e da matriz, fechar diferencas de payload, ownership, paginacao e eventos. Aceite com fixtures de protocolo e duas contas de teste.
2. Midia e integracoes: persistencia automatica e retencao, Chatwoot/n8n, Redis Streams/SQS; testar indisponibilidade, recuperacao, duplicatas e limites antes de publicar cada adaptador.
3. Comercio: leitura de catalogo/produto antes de mutacoes e pedidos. Preservar moeda/unidade, diferenciar aceite de protocolo de pagamento confirmado e usar idempotencia nas mutacoes suportadas.
4. Interativos: validar botoes/listas, depois carrosseis/flows, com fixtures e recebimento observado em Android, iOS e Web. Nao converter resposta HTTP aceita em paridade comprovada.
5. Chamadas: separar rejeicao/sinalizacao de transporte de audio/video e chamadas de grupo. Exigir autenticacao de streams, limites, limpeza, codec compativel e teste real de ponta a ponta.
6. Carga e restauracao: medir CPU/RSS/disco, p95/p99, backlog e reconexao com perfis reproduziveis; restaurar backup consistente e repetir jornadas antes do aceite final.

Cada jornada deve atualizar a matriz com SHA, contrato, teste, resultado e lacuna restante. As 303 operacoes do baseline continuam identificadas como baseline, sem substituir evidencia funcional por uma contagem de rotas novas.

## Pareamento real - QR expirado
- [x] Usuario informou "QR invalido ou expirado"; diagnostico confirmou instancia saudavel e ciclo anterior encerrado sem CONNECTED.
- [x] Substituida imagem estatica do chat por pagina HTTPS temporaria, caminho privado aleatorio, validade de 30 minutos, sem cache e sem expor credencial do Manager.
- [x] Verificados PNG atual e mudanca real da versao do QR. A pagina consulta o estado a cada segundo e reinicia o ciclo encerrado.
- [x] Validado JavaScript efetivamente renderizado; erro de escape identificado e corrigido antes de entregar o link. Teste test-live-pairing.cjs aprovado para rotacao, ocultacao de QR antigo, estado conectado e expiracao.
- [x] Confirmar CONNECTED depois do escaneamento: servidor confirmou conexao em 10/11/2026 19:19 UTC; usuario confirmou sucesso.
- [x] Executar envio/recebimento reais autorizados. Primeiro envio retornou 502 e foi preservado no ledger sem reenvio automatico; apos a correcao, novo teste identificado de texto e arquivo recebeu HTTP 201, ACK de entrega/leitura e resposta real.
- [x] Removidos helper HTTPS, rota e configuracao privada depois de confirmar CONNECTED. Timer de limpeza desativado; sessao de teste preservada. Recibo live-pairing-cleanup.json.
- [x] Agendada limpeza automatica do helper expirado via timer systemd dedicado; limpeza antecipada apos pareamento concluida e timer desativado. Script verifica nome, label e diretorio antes de remover somente o helper e sua configuracao.

## Correcao encontrada no aceite real
- [x] Isolar falha de resolucao: probe sem envio confirmou GetUserInfo com perfil valido sem LID, ResolveLID recusando destinatario e IsOnWhatsApp retornando contato registrado em LID. Conta de teste diferente do destinatario; nenhum dado privado publicado.
- [x] SDK: fallback para consulta de contatos quando perfil omite LID; exige correspondencia da consulta e um unico contato registrado em LID. Preserva device e mappings canonicos.
- [x] Regressao reproduzida antes da correcao por serializacao e parsing reais de USync. Seis cenarios, suite completa e vet aprovados; CI Go antigo/atual, race, WASM, MEX e pre-commit aprovados na run 42886170997.
- [x] HyperMeow PR #4 incorporado: d326350a86e98c405c9f6adf14cb6bd03ea8e66a. Fonte da correcao 44b882797ce70f44ce41d1d31d7461824d92a6c0.
- [x] Atualizar dependencia do server-go, testar, incorporar e publicar imagem validada; repetir aceite real e renovar reservas. Server-go PR #20, main e01698bf3e8b0b9d9dc749e7ffbb49e65deab22f; CI 42887160468 aprovado.

## Aceite real e pool atualizada - 2026-11-10
- Imagem: sha256:b0d96f63e775038ff4c71cfdf4630763712fe8e978ff555dedb685cae8ed9b5b. Arquivo-fonte da main limpa: SHA256 17c5ed3a0018f94f2b3527b547ccc751b90092060067d62e402d43cb454e735e.
- Canary isolado repetiu recuperacao de webhook, auditoria, isolamento, ledger e restart. Fonte conferida novamente antes do rollout. Evidencias recipient-fix-built.json e recipient-fix-canary.json.
- Primeiro ensaio de troca da conta de teste falhou por alias configurado na rede bridge padrao. Script corrigido; imagem anterior restaurada com o volume preservado e CONNECTED confirmado. Segunda troca concluida com backup consistente, mesma porta/configuracao e pareamento preservado. Nao houve alteracao de sessoes de clientes.
- Texto e arquivo entregues e lidos pelo destinatario autorizado; resposta exata de teste recebida pelo servidor. Repeticao da mesma chave recuperou resultado, alteracao de payload retornou 409, outro workspace nao acessou o resultado. Recibo publico sem identificadores: phone-acceptance-summary.json.
- Restart preservou pareamento, os dois resultados bem-sucedidos e a falha inicial. Download do arquivo depois do restart identico byte a byte. Backup restaurado offline com integridade dos cinco bancos SQLite validada. Recibo phone-restart-validation.json.
- Tres reservas anteriores renovadas por alocacoes de teste proprias e exclusao pelo Manager. Todas usam a nova imagem; Manager e worker saudaveis, sessao anterior preservada. Recibos recipient-fix-release.json, recipient-fix-live-pool.json e recipient-fix-final-audit.json.
- Helper de pareamento, rota HTTPS, configuracao privada e container do probe removidos. Sessao Go de teste permanece pareada para as jornadas restantes. Fonte da imagem privada, logout, segunda conta, alertas, carga e paridade integral continuam com as limitacoes registradas acima.


## Canais e integracoes - 2026-11-10

- [HyperMeow #6](https://github.com/wppconnect-team/hypermeow/pull/6): cancelamento, desconexao e ACK negativo nao produzem sucesso de visualizacao. CI 42893049289; main 099af656ac43dcd4b3e2688f8a880ee76e2c4e75.
- [Server Go #21](https://github.com/wppconnect-team/wppconnect-server-go/pull/21): oito operacoes, seis mutacoes com ledger persistente e duas consultas separadas de resultado. Testes de contrato HTTP, isolamento, permissao e replay; CI 42895236212 com race, integracoes e imagens amd64/arm64. Main 4fb3d7f13de475af0e6fa7bcd89b733b12105121. [Guia de canais](https://github.com/wppconnect-team/wppconnect-server-go/blob/main/docs/newsletters.pt-BR.md).
- [Cloud #40](https://github.com/joaosouz4dev/wppconnect-cloud/pull/40): 49 testes, controle de escopos, typecheck, build e CI 42895717043. Chromium: guia em ingles (1440 px) e portugues (375 px), link do contrato, foco de teclado e largura da secao conferidos.
- [Server Go #22](https://github.com/wppconnect-team/wppconnect-server-go/pull/22): adaptador Redis Streams opcional, fixture de Redis real no CI, inspecao e replay por destino. Nenhuma fila ou stream de cliente configurado.
- A matriz original continua sendo o retrato de 303 operacoes da versao comparada. Nomes e payloads sao nativos; Z025 nao tem filtro before no SDK. A reacao recebe 202 sem ACK remoto e um convite administrativo nao concede o papel.

- Canary do binario e rollout das tres reservas Go aprovados: SHA da imagem 59f610bd5d2060cd4b1eecff4ac134f7a41a49d696e97bdac68461b4eeff508a. Auditoria 20:41 UTC: tres reservas healthy, containers temporarios removidos e nenhum OOM, panic, fatal ou falha de persistencia nos logs verificados. A conta de testes pareada e a alocacao preexistente foram preservadas.
- [Server Go #23](https://github.com/wppconnect-team/wppconnect-server-go/pull/23): leitura do catalogo proprio, produtos e resumos de colecoes. Testes locais de escopo, preco/moeda, visibilidade falsa, cursor, limites e erros aprovados.


## Brokers, comercio e carga - 2026-11-10

- Redis Streams: main ba6dee6414e7630e18ced46c0be12b74fbf69b62, CI 42896819299. O ID de evento permanece estavel mesmo quando uma entrega repetida gera outro ID de entrada no stream.
- Catalogo proprio: main 1c42366e18125015a9d30d7abf3a5cafedeccb31, CI 42898388266. A Cloud usa commerce:read em tres consultas. Testes de conta Business real, mutacoes, pedidos e pagamentos continuam abertos.
- SQS: main 666afa223c12efefcb9ac90a2d6920d9ea3a1b11, CI 42898825516. O teste de integracao usa ElasticMQ, nao uma conta AWS. Tanto SQS como Redis permanecem desativados por padrao.
- [Carrossel experimental #25](https://github.com/wppconnect-team/wppconnect-server-go/pull/25): validacao local HTTP/SQLite/SDK, rejeicao antes de upload e resposta idempotente. Nao enviar para validar clientes fisicos enquanto esse teste estiver adiado.
- Benchmark isolado: limite 0,75 CPU e 256 MiB, 24 consultas sequenciais e 48 com quatro clientes. P95 de 98,50 ms e 575,48 ms respectivamente. Primeiro ensaio nao concluiu o restart por reutilizar a porta efemera anterior; o ensaio completo consultou novamente a porta e recuperou os quatro registros. Sem mensagens de WhatsApp. [Medicoes](docs/benchmark-result-lookup.json) e [script](docs/benchmark-result-lookup.py).

- Publicacao de brokers/catalogo: imagem 49a88afea20f64521ed3bb453aeb4ef3bbc77ba45b43d011abd0645cabf8db78. Tres reservas Go healthy, sem OOM ou erros criticos recentes. Cloud #41, main 3751275b5385c92ea25f5db623669ac97aad4baa, deploy confirmado e guia publico conferido em dois idiomas.

- CI da fonte combinada brokers/catalogo: [42900118182](https://github.com/wppconnect-team/wppconnect-server-go/actions/runs/42900118182). O OpenAPI no Git era identico byte a byte ao gerado; o primeiro checkout Windows materializou CRLF. Repeticao com os bytes do Git passou e a regra de LF foi incorporada no PR #25.
