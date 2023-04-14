package cepv1

type CepOutput struct {
	Cep          string
	State        string
	City         string
	Neighborhood string
	Street       string
	Service      string
}

func (c Client) FindByCepWithMultiplesProviders(cep string) (*CepOutput, error) {
    var output CepOutput
    
    if err := c.handler(cep, &output); err != nil {
        return nil, err 
    }
	return &output, nil
}

