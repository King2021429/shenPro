package model

// StockQueryReq 商业化 库存查询请求结构体
type StockQueryReq struct {
	SkuIdList []int64 `json:"sku_id_list" form:"sku_id_list"`
}

// StockInfo 商业化 库存查询返回 结构体
type StockInfo struct {
	IsolateStock  int64 `json:"isolate_stock"`
	PerformsStock int64 `json:"performs_stock"`
	SkuId         int64 `json:"sku_id"`
	Stock         int64 `json:"stock"`
}

// StockUpdateReq 商业化 库存更新请求结构体
type StockUpdateReq struct {
	Uid             int64              `json:"uid"`
	ClientId        string             `json:"client_id"`
	Mode            int64              `json:"mode"`
	ModifyStockList []*ModifyStockInfo `json:"modify_stock_list"`
}

type ModifyStockInfo struct {
	SkuId               int64 `json:"sku_id"`
	ZpEntityStockOffset int64 `json:"zp_entity_stock_offset"`
}

// StockUpdateResp 商业化 库存更新返回 结构体
type StockUpdateResp struct {
	FailedSkus []*FailedSku `json:"failed_skus"`
}

type FailedSku struct {
	StockId int64  `json:"stock_id"`
	Message string `json:"message"`
}
