# 项目分层设计概述（基于 Gorm 驱动的数据库操作）

本设计总结了典型的分层架构，包括 API 层、Handler 层、Logic 层、Use Case（UC）层、Repository 层和 Model 层。数据库操作通过 Gorm 驱动实现，以下为详细设计说明和层次划分：

---

# **1. 分层架构概览**

1. **API 层**：
   - 描述：提供与外部通信的接口层，定义 HTTP 请求路径和数据结构。
   - 作用：接收请求、返回响应，并调用对应的 Handler 层。
   - 示例文件路径：`api/`

2. **Handler 层**：
   - 描述：处理 API 请求的逻辑分发层。
   - 作用：解析请求参数，调用 Logic 层进行业务处理，返回结果。
   - 示例文件路径：`internal/handler/`

3. **Logic 层**：
   - 描述：核心业务流程管理层。
   - 作用：协调多个 Use Case 层的调用，管理事务，处理业务流程中的异常。
   - 示例文件路径：`internal/logic/`

4. **Use Case 层（UC 层）**：
   - 描述：实现具体的业务逻辑细节。
   - 作用：完成单一功能逻辑（如库存检查、订单创建等），对 Repository 层调用的封装。
   - 示例文件路径：`internal/uc/`

5. **Repository 层**：
   - 描述：数据访问层，封装对数据库的具体操作。
   - 作用：提供对数据库的 CRUD 操作接口，隔离底层实现细节。
   - 示例文件路径：`internal/repository/`

6. **Model 层**：
   - 描述：定义与数据库表结构对应的模型。
   - 作用：为 Repository 层提供基础的数据库映射结构。
   - 示例文件路径：`internal/model/`

---

## **2. 各层职责和实现细节**

### **API 层**
- 主要功能：
  - 使用 .api 文件定义接口，自动生成 handler 和数据模型。
  - 通过 API 定义服务的路由和输入输出数据结构。
- 示例代码：
```go
// api/order.api
// api/order.api
type (
    CreateOrderReq {
        userId int64
        items []OrderItem
    }

    CreateOrderResp {
        orderId int64
    }

    OrderItem {
        productId int64
        quantity  int
    }
)

@handler CreateOrderHandler
post /orders (CreateOrderReq) returns (CreateOrderResp)

```

---

#### **Handler 层**
- 主要功能：
  - 解析 API 层传入的请求数据。
  - 调用 Logic 层完成业务流程。
  - 统一返回响应数据。
- 示例代码：
```go
// internal/handler/order_handler.go
type OrderHandler struct {
    logic *logic.OrderLogic
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var req CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.logic.CreateOrder(c.Request.Context(), req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
```

---

### **Logic 层**
- 主要功能：
  - 组织和协调业务流程。
  - 管理事务的开启、提交和回滚。
  - 调用 UC 层处理具体的业务逻辑。
- 示例代码：
```go
// internal/logic/order_logic.go
type OrderLogic struct {
    db       *gorm.DB
    useCases *usecase.OrderUseCases
}

func (l *OrderLogic) CreateOrder(ctx context.Context, req CreateOrderRequest) error {
    // 开启事务
    return l.db.Transaction(func(tx *gorm.DB) error {
        // 调用 Use Case 层的功能
        if err := l.useCases.CheckStock(ctx, tx, req.Items); err != nil {
            return err
        }
        return l.useCases.CreateOrder(ctx, tx, req)
    })
}
```

---

#### **Use Case 层**
- 主要功能：
  - 实现单一职责的业务逻辑。
  - 封装对 Repository 层的调用，完成独立的业务功能。
- 示例代码：
```go
// internal/logic/usecase/order_usecase.go
type OrderUseCases struct {
    repo *repository.OrderRepository
}

func (uc *OrderUseCases) CheckStock(ctx context.Context, tx *gorm.DB, items []OrderItem) error {
    for _, item := range items {
        stock, err := uc.repo.GetStock(ctx, tx, item.ProductID)
        if err != nil {
            return err
        }
        if stock < item.Quantity {
            return fmt.Errorf("insufficient stock for product %d", item.ProductID)
        }
    }
    return nil
}

func (uc *OrderUseCases) CreateOrder(ctx context.Context, tx *gorm.DB, req CreateOrderRequest) error {
    return uc.repo.CreateOrder(ctx, tx, req)
}
```

---

##### **Repository 层**
- 主要功能：
  - 封装对数据库的具体操作，提供接口供 UC 层调用。
  - 直接使用 Gorm 完成对数据库的读写操作。
- 示例代码：
```go
// internal/repository/order_repository.go
type OrderRepository struct {
    db *gorm.DB
}

func (r *OrderRepository) GetStock(ctx context.Context, tx *gorm.DB, productID int) (int, error) {
    var product model.Product
    if err := tx.WithContext(ctx).First(&product, productID).Error; err != nil {
        return 0, err
    }
    return product.Stock, nil
}

func (r *OrderRepository) CreateOrder(ctx context.Context, tx *gorm.DB, req CreateOrderRequest) error {
    order := model.Order{
        UserID: req.UserID,
        Items:  req.Items,
    }
    return tx.WithContext(ctx).Create(&order).Error
}
```

---

##### **Model 层**
- 主要功能：
  - 定义数据库表的结构体映射。
  - 提供 Gorm 所需的基础结构和标签。
- 示例代码：
```go
// internal/model/order.go
type Order struct {
    ID     uint           `gorm:"primaryKey"`
    UserID int            `gorm:"not null"`
    Items  []OrderItem    `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    ID        uint `gorm:"primaryKey"`
    OrderID   uint
    ProductID int  `gorm:"not null"`
    Quantity  int  `gorm:"not null"`
}
```

---

#### **3. 总结**

- **层次划分清晰**：API、Handler、Logic、UC、Repository 和 Model 各司其职，分工明确。
- **事务管理在 Logic 层**：通过 Gorm 的事务管理（`Transaction` 方法）实现事务的统一控制。
- **单一职责原则**：每个层只负责自己的核心职责，避免逻辑耦合。
- **易于扩展**：新需求可以通过添加新的 UC 或 Repository 轻松实现。
- **适配 Gorm**：Repository 层直接封装 Gorm 操作，与业务逻辑解耦。

这套设计既保持了代码结构的清晰性，又便于维护和扩展，可作为项目的标准参考架构。