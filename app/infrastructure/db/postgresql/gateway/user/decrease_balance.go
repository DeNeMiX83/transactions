package user

import (

)

func (g *gateway) DecreaseBalance (user_id string, balanc_id string, amount int) error{
	sqlStatement := `
	UPDATE balance
	SET amount = amount - $1
	WHERE user_id = $2 and id = $3;`
	_, err := g.db.Exec(sqlStatement, amount, user_id, balanc_id)
	if err != nil {
		return err
	}
	return nil
}