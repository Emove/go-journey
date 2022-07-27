package facade

import "fmt"

// OrderFacade 假设为订单模块的外观接口
type OrderFacade interface {
	Order()
}

// AccountModuleApi 假设为账户信息Api接口
type AccountModuleApi struct {
}

// 获取账户余额信息
func (api AccountModuleApi) GetAccountBalanceInfo() {
	fmt.Println("查询账户余额信息")
}

func (api AccountModuleApi) Deduction() {
	fmt.Println("扣除账户余额")
}

type OrderFacadeImpl struct {
	accountApi *AccountModuleApi
}

func NewOrderFacade() OrderFacade {
	return &OrderFacadeImpl{
		accountApi: &AccountModuleApi{},
	}
}

func (facade *OrderFacadeImpl) Order() {
	accountApi := facade.accountApi
	accountApi.GetAccountBalanceInfo()
	accountApi.Deduction()
	fmt.Println("order success")
}
