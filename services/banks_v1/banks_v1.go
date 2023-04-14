package banksv1


type BanksOutput struct {
    Ispb string
    Name string
    Code int
    FullName string
}

func (c Client) FindBank(code string) (*BanksOutput, error ){
    var output BanksOutput
    if err := c.handler(code, &output); err != nil {
        return nil, err
    }
    return &output, nil 
}

func (c Client) FindBanks() (*[]BanksOutput, error ){
    var outputs []BanksOutput
    if err := c.handler("", &outputs); err != nil {
        return nil, err
    }
    return &outputs, nil
}
