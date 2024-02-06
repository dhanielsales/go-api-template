package entity

type NoDb struct { // TODO Fazer implementação de outro repository a partir da mesma assinatura
	ID   string `json:"id"`
	Name string `json:"name"`
}
