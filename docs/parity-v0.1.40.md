# Paridade isolada com a Zuck v2.4.2 na v0.1.40

A release de referência é a [v0.1.40 do server-go](https://github.com/wppconnect-team/wppconnect-server-go/releases/tag/v0.1.40), commit `0b4b3091377a42280145df134d1625580f00312c`. O manifesto OCI usa `sha256:b9a33faca0a1c8cfef7f6eab7b6d33324f64b56c04acefc52e5cc3cda401abe3` e inclui `linux/amd64` e `linux/arm64`. As verificações desta rodada rodaram com contas, dados e destinos sintéticos em processos e contêineres descartáveis.

## Comparação de contratos

- [x] A matriz congelada tem 303 operações da Zuck. O OpenAPI da v0.1.40 tem 170 operações e o Swagger do Node fixado em `37b7bc1a662b5262b8276597701f38c4dbcad58a` tem 158. A interseção literal por método e caminho é 87, com 83 operações exclusivas do Go e 71 do Node. A contagem descreve a superfície HTTP, incluindo métodos experimentais ou indisponíveis declarados no contrato.
- [x] Sete funções passaram para os 22 contratos nativos mapeados. Os dez diferenciais antes identificados entre o Node e o Go têm contratos Go testados. [Matriz](parity-contract-matrix.md) · [Mapeamento](parity-progress.json)
- [x] Z290, Z271, Z282, Z285 e Z283 passaram nos testes de exclusão de canal e coleção, criação, edição e exclusão de produtos. As fixtures exercitam ID e papel exigidos para excluir canal, limites de lote, precisão de preço, moeda, revisão do produto, visibilidade, contagem parcial e recuperação idempotente.
- [x] Um processo Go desconectado foi encerrado e recriado com o mesmo SQLite. Os cinco resultados HTTP 404 originais continuaram consultáveis. A repetição devolveu o mesmo corpo e o cabeçalho de replay. A mesma chave com outro corpo recebeu 409. [Teste](check-isolated-mutation-replay.py)
- [ ] Executar a comparação HTTP com a Zuck v2.4.2 licenciada. O Swagger permite comparar campos e códigos declarados. Os corpos reais, erros operacionais, efeitos remotos, recuperações e limites exigem executar ambos os servidores. A referência encerrou a inicialização nos ensaios anteriores sem `LICENSE_KEY`. [Registro da referência](reference-runtime-acceptance.json)

## Sessões e restauração

- [x] Um contêiner descartável da imagem publicada verificou `pair-phone` sem credencial (401), com token de outra sessão (401) e com telefone inválido (400). O logout desconectado retornou 409. `close-session` retornou 200 e o estado `CLOSED` permaneceu após reiniciar.
- [x] As fixtures do Manager testaram validação, cancelamento e geração de código pelo cliente simulado. Testaram também preservação de credenciais quando o logout remoto falha e revogação local quando tem sucesso. A migração SQLite da v16 à v21 verificou chaves, segredo WASA, registro histórico e repetição do upgrade. [Ciclo de sessão](https://github.com/wppconnect-team/wppconnect-server-go/blob/0b4b3091377a42280145df134d1625580f00312c/internal/session/lifecycle_test.go) · [Pareamento](https://github.com/wppconnect-team/wppconnect-server-go/blob/0b4b3091377a42280145df134d1625580f00312c/internal/session/settings_test.go)
- [ ] Confirmar vínculo por código e revogação remota em conta descartável. O endpoint de código aguarda a prontidão do transporte inicial. Nenhum código foi solicitado a um telefone real nesta rodada.

## Suites e serviços isolados

- [x] No Windows, 373 testes e subtestes do server-go e 793 do SDK passaram, junto com vet e builds. O SDK teve também verificação MEX e build JS/WASM. O checkout do SDK tinha o mesmo código e migrações da revisão usada na tag.
- [x] No Docker local, seis testes integraram Redis Streams, ElasticMQ, NATS, RabbitMQ e MinIO. Cobriram recuperação de backlog e leases após reinício, retenção Redis, rejeição e deduplicação SQS FIFO, reconexão NATS/RabbitMQ e integridade e limpeza no S3 privado. A rede, os contêineres e o volume da fixture foram removidos.
- [x] O CI de publicação da v0.1.40 passou com detector de corridas, Redis, ElasticMQ, NATS, RabbitMQ, MinIO, n8n e smoke em amd64 e arm64. O teste n8n importou o workflow de exemplo, recusou a credencial inválida e processou um evento sintético autenticado. [CI da tag](https://github.com/wppconnect-team/wppconnect-server-go/actions/runs/46039283234)
- [x] O smoke local da imagem publicada passou em amd64 e arm64 (emulado), com inicialização, prontidão, encerramento e reinício. O estado fechado persistiu. No Docker Desktop, o arm64 foi baixado pelo digest de seu manifesto filho para evitar a colisão do índice já associado ao amd64.

- [x] Em novo CI da tag v0.1.40, o fixture Chatwoot 4.19.1 assinou um callback que o adaptador aceitou, persistiu uma mensagem de texto após entregas repetidas e recuperou as referências do provisionamento depois de respostas interrompidas. A inbox, o PostgreSQL e o Redis eram descartáveis, e a ponta WhatsApp foi representada por HTTP local. [CI Chatwoot](https://github.com/wppconnect-team/wppconnect-server-go/actions/runs/46069472875)

O n8n e o Chatwoot foram confirmados por CI na própria tag. As réplicas locais dessas duas integrações ficaram no download das imagens e foram encerradas após a confirmação no CI.

## Aceites pendentes

| Dependência | Critério de encerramento |
| --- | --- |
| WhatsApp Business com catálogo de teste | Ler e paginar produtos e coleções, validar efeitos de criação, edição, exclusão e visibilidade, conferir revisão, leitura no comprador e estado do carrinho. |
| Instalação licenciada da Zuck v2.4.2 | Executar a matriz comportamental lado a lado. Validar corpos, falhas, repetição, isolamento, integração e efeitos remotos. |
| Decisões de teste já registradas | Exclusão real de canal segue restrita a testes isolados. Segunda conta e clientes Android, iPhone e Web continuam adiados. Monitoramento externo continua pendente de serviço e destino. |
| Implementação e testes restantes | Anexos bidirecionais no Chatwoot, arquivamento automático com opt-in, pedidos e pagamentos, entrega de Flows e interativos, áudio/vídeo e chamadas em grupo, carga de protocolo com metas. |

A confirmação de texto do Chatwoot e o upload explícito com retenção em S3 cobrem partes das duas primeiras integrações. Os critérios completos para lançamentos seguintes estão em [próximas entregas](parity-remaining.md).
