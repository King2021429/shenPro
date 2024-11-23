package model

type DeveloperAppDlcOrderInfo struct {
	// OrderId 订单id
	OrderId string `json:"order_id"`
	// dlc id
	DlcId string `json:"dlc_id"`
	// 订单状态 processing:待发货流程中 success:订单成功 fail:订单失败已退款
	OrderStatus int64 `json:"order_status"`
	// 订单时间
	OrderCreateTime int64 `json:"order_create_time"`
	// 成功失败时间
	OrderFinishTime int64 `json:"order_finish_time"`
}

type QueryReq struct {
	// dlc_id 必传
	DlcId string `json:"dlc_id"`
	// 订单id 必传
	OrderIds []string `json:"order_ids"`
}

type QueryResp struct {
	OrderList []DeveloperAppDlcOrderInfo `json:"order_list"`
}
