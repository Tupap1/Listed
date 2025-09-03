package database


import (
	"log"
	"github.com/Tupap1/Listed/server/models"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	log.Println("🌱 Iniciando seeding de la base de datos...")

	if err := SeedRoles(db); err != nil {
		return err
	}

	if err := SeedPermissions(db); err != nil {
		return err
	}

	if err := SeedRolePermissions(db); err != nil {
		return err
	}

	if err := SeedDefaultUser(db); err != nil {
		return err
	}

	log.Println("✅ Seeding completado exitosamente!")
	return nil
}

func SeedRoles(db *gorm.DB) error {
	log.Println("📝 Creando roles...")

	roles := []models.Role{
		{Name: "owner", Description: "Propietario del negocio - acceso completo", IsDefault: false},
		{Name: "admin", Description: "Administrador del sistema", IsDefault: false},
		{Name: "manager", Description: "Gerente de inventario", IsDefault: false},
		{Name: "warehouse_keeper", Description: "Encargado de bodega", IsDefault: false},
		{Name: "seller", Description: "Vendedor", IsDefault: true},
		{Name: "viewer", Description: "Solo consulta", IsDefault: false},
	}

	for _, role := range roles {
		if err := db.Where("name = ?", role.Name).FirstOrCreate(&role).Error; err != nil {
			log.Printf("❌ Error creando rol %s: %v", role.Name, err)
			return err
		}
		log.Printf("✓ Rol creado/verificado: %s", role.Name)
	}

	return nil
}

func SeedPermissions(db *gorm.DB) error {
	log.Println("🔐 Creando permisos...")

	permissions := []models.Permission{
		// Productos
		{Name: "products.read", Description: "Ver productos", Resource: "products", Action: "read"},
		{Name: "products.create", Description: "Crear productos", Resource: "products", Action: "create"},
		{Name: "products.update", Description: "Modificar productos", Resource: "products", Action: "update"},
		{Name: "products.delete", Description: "Eliminar productos", Resource: "products", Action: "delete"},
		{Name: "products.import", Description: "Importar productos masivamente", Resource: "products", Action: "import"},

		{Name: "inventory.read", Description: "Ver inventario", Resource: "inventory", Action: "read"},
		{Name: "inventory.update", Description: "Ajustar inventario", Resource: "inventory", Action: "update"},
		{Name: "inventory.transfer", Description: "Transferir entre bodegas", Resource: "inventory", Action: "transfer"},
		{Name: "inventory.audit", Description: "Hacer auditorías de inventario", Resource: "inventory", Action: "audit"},

		// Ventas
		{Name: "sales.read", Description: "Ver ventas", Resource: "sales", Action: "read"},
		{Name: "sales.create", Description: "Crear ventas", Resource: "sales", Action: "create"},
		{Name: "sales.update", Description: "Modificar ventas", Resource: "sales", Action: "update"},
		{Name: "sales.cancel", Description: "Cancelar ventas", Resource: "sales", Action: "cancel"},
		{Name: "sales.refund", Description: "Hacer devoluciones", Resource: "sales", Action: "refund"},

		// Compras
		{Name: "purchases.read", Description: "Ver compras", Resource: "purchases", Action: "read"},
		{Name: "purchases.create", Description: "Crear órdenes de compra", Resource: "purchases", Action: "create"},
		{Name: "purchases.approve", Description: "Aprobar compras", Resource: "purchases", Action: "approve"},
		{Name: "purchases.receive", Description: "Recibir mercancía", Resource: "purchases", Action: "receive"},

		// Proveedores y Clientes
		{Name: "suppliers.read", Description: "Ver proveedores", Resource: "suppliers", Action: "read"},
		{Name: "suppliers.create", Description: "Crear proveedores", Resource: "suppliers", Action: "create"},
		{Name: "suppliers.update", Description: "Modificar proveedores", Resource: "suppliers", Action: "update"},
		{Name: "customers.read", Description: "Ver clientes", Resource: "customers", Action: "read"},
		{Name: "customers.create", Description: "Crear clientes", Resource: "customers", Action: "create"},
		{Name: "customers.update", Description: "Modificar clientes", Resource: "customers", Action: "update"},

		// Reportes
		{Name: "reports.sales", Description: "Ver reportes de ventas", Resource: "reports", Action: "sales"},
		{Name: "reports.inventory", Description: "Ver reportes de inventario", Resource: "reports", Action: "inventory"},
		{Name: "reports.financial", Description: "Ver reportes financieros", Resource: "reports", Action: "financial"},
		{Name: "reports.export", Description: "Exportar reportes", Resource: "reports", Action: "export"},

		// Configuraciones
		{Name: "settings.read", Description: "Ver configuraciones", Resource: "settings", Action: "read"},
		{Name: "settings.update", Description: "Modificar configuraciones", Resource: "settings", Action: "update"},
		{Name: "settings.system", Description: "Configuraciones del sistema", Resource: "settings", Action: "system"},

		// Usuarios
		{Name: "users.read", Description: "Ver usuarios", Resource: "users", Action: "read"},
		{Name: "users.create", Description: "Crear usuarios", Resource: "users", Action: "create"},
		{Name: "users.update", Description: "Modificar usuarios", Resource: "users", Action: "update"},
		{Name: "users.delete", Description: "Eliminar usuarios", Resource: "users", Action: "delete"},

		// Bodegas
		{Name: "warehouses.read", Description: "Ver bodegas", Resource: "warehouses", Action: "read"},
		{Name: "warehouses.create", Description: "Crear bodegas", Resource: "warehouses", Action: "create"},
		{Name: "warehouses.update", Description: "Modificar bodegas", Resource: "warehouses", Action: "update"},
	}

	for _, permission := range permissions {
		if err := db.Where("name = ?", permission.Name).FirstOrCreate(&permission).Error; err != nil {
			log.Printf("❌ Error creando permiso %s: %v", permission.Name, err)
			return err
		}
	}

	log.Printf("✓ %d permisos creados/verificados", len(permissions))
	return nil
}

// SeedRolePermissions - Asigna permisos a roles
func SeedRolePermissions(db *gorm.DB) error {
	log.Println("🔗 Asignando permisos a roles...")

	// Definir qué permisos tiene cada rol
	rolePermissions := map[string][]string{
		"owner": {
			// Owner tiene TODOS los permisos
			"products.read", "products.create", "products.update", "products.delete", "products.import",
			"inventory.read", "inventory.update", "inventory.transfer", "inventory.audit",
			"sales.read", "sales.create", "sales.update", "sales.cancel", "sales.refund",
			"purchases.read", "purchases.create", "purchases.approve", "purchases.receive",
			"suppliers.read", "suppliers.create", "suppliers.update",
			"customers.read", "customers.create", "customers.update",
			"reports.sales", "reports.inventory", "reports.financial", "reports.export",
			"settings.read", "settings.update", "settings.system",
			"users.read", "users.create", "users.update", "users.delete",
			"warehouses.read", "warehouses.create", "warehouses.update",
		},
		"admin": {
			"products.read", "products.create", "products.update", "products.delete", "products.import",
			"inventory.read", "inventory.update", "inventory.transfer", "inventory.audit",
			"sales.read", "sales.create", "sales.update", "sales.cancel", "sales.refund",
			"purchases.read", "purchases.create", "purchases.approve", "purchases.receive",
			"suppliers.read", "suppliers.create", "suppliers.update",
			"customers.read", "customers.create", "customers.update",
			"reports.sales", "reports.inventory", "reports.financial", "reports.export",
			"settings.read", "settings.update", // NO settings.system
			"users.read", "users.create", "users.update", // NO users.delete
			"warehouses.read", "warehouses.create", "warehouses.update",
		},
		"manager": {
			"products.read", "products.create", "products.update",
			"inventory.read", "inventory.update", "inventory.transfer", "inventory.audit",
			"sales.read",
			"purchases.read", "purchases.create", "purchases.approve",
			"suppliers.read", "suppliers.create", "suppliers.update",
			"customers.read",
			"reports.sales", "reports.inventory", "reports.export",
			"warehouses.read",
		},
		"warehouse_keeper": {
			"products.read",
			"inventory.read", "inventory.update", "inventory.transfer",
			"purchases.read", "purchases.receive",
			"suppliers.read",
			"warehouses.read",
		},
		"seller": {
			"products.read",
			"inventory.read",
			"sales.read", "sales.create",
			"customers.read", "customers.create", "customers.update",
		},
		"viewer": {
			"products.read",
			"inventory.read",
			"sales.read",
			"customers.read",
			"suppliers.read",
			"warehouses.read",
		},
	}

	for roleName, permissionNames := range rolePermissions {
		// Buscar el rol
		var role models.Role
		if err := db.Where("name = ?", roleName).First(&role).Error; err != nil {
			log.Printf("❌ No se encontró el rol: %s", roleName)
			continue
		}

		// Limpiar permisos existentes del rol
		if err := db.Model(&role).Association("Permissions").Clear(); err != nil {
			log.Printf("❌ Error limpiando permisos del rol %s: %v", roleName, err)
			continue
		}

		// Asignar nuevos permisos
		for _, permName := range permissionNames {
			var permission models.Permission
			if err := db.Where("name = ?", permName).First(&permission).Error; err != nil {
				log.Printf("❌ No se encontró el permiso: %s", permName)
				continue
			}

			if err := db.Model(&role).Association("Permissions").Append(&permission); err != nil {
				log.Printf("❌ Error asignando permiso %s al rol %s: %v", permName, roleName, err)
			}
		}

		log.Printf("✓ Permisos asignados al rol: %s (%d permisos)", roleName, len(permissionNames))
	}

	return nil
}



// SeedDefaultUser - Crea el usuario administrador inicial
func SeedDefaultUser(db *gorm.DB) error {
	log.Println("👤 Creando usuario administrador por defecto...")

	// Verificar si ya existe un usuario owner
	var existingUser models.User
	err := db.Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.name = ?", "owner").
		First(&existingUser).Error

	if err == nil {
		log.Println("✓ Ya existe un usuario owner, omitiendo creación")
		return nil
	}

	// Crear usuario owner
	user, err := models.CreateUser(db, "admin", "admin@tuempresa.com", "admin123", nil)
	if err != nil {
		log.Printf("❌ Error creando usuario por defecto: %v", err)
		return err
	}

	// Asignar rol owner
	var ownerRole models.Role
	if err := db.Where("name = ?", "owner").First(&ownerRole).Error; err != nil {
		log.Printf("❌ No se encontró el rol owner: %v", err)
		return err
	}

	if err := db.Model(user).Association("Roles").Append(&ownerRole); err != nil {
		log.Printf("❌ Error asignando rol owner: %v", err)
		return err
	}

	log.Println("✅ Usuario administrador creado:")
	log.Println("   Email: admin@tuempresa.com")
	log.Println("   Password: admin123")
	log.Println("   ⚠️  CAMBIAR PASSWORD EN PRODUCCIÓN")

	return nil
}