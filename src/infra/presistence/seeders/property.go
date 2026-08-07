package seeders

import (
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"gorm.io/gorm"
)

var defaultPropertyCategories []models.PropertyCategory = []models.PropertyCategory{
	{Name: "Body", Icon: "Body"},
	{Name: "Engine", Icon: "Engine"},
	{Name: "Drivetrain", Icon: "Drivetrain"},
	{Name: "Suspension", Icon: "Suspension"},
	{Name: "Equipment", Icon: "Equipment"},
	{Name: "Driver support systems", Icon: "Driver support systems"},
	{Name: "Lights", Icon: "Lights"},
	{Name: "Multimedia", Icon: "Multimedia"},
	{Name: "Safety equipment", Icon: "Safety equipment"},
	{Name: "Seats and steering wheel", Icon: "Seats and steering wheel"},
	{Name: "Windows and mirrors", Icon: "Windows and mirrors"},
}

func SeedProperties(database *gorm.DB) error {
	count := 0

	database.
		Model(&models.PropertyCategory{}).
		Select("COUNT(*)").
		Find(&count)

	for _, v := range defaultPropertyCategories {
		if err := database.Model(&models.PropertyCategory{}).Create(&v).Error; err != nil {
			return err
		}
		createProperty(database, v.Name)
	}

	return nil
}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.
		Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select("COUNT(*)").
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}
	var props *[]models.Property
	switch cat {
	case "Body":
		props = getBodyProperties(int(catModel.Id))

	case "Engine":
		props = getEngineProperties(int(catModel.Id))

	case "Drivetrain":
		props = getDrivetrainProperties(int(catModel.Id))

	case "Suspension":
		props = getSuspensionProperties(int(catModel.Id))

	case "Comfort":
		props = getComfortProperties(int(catModel.Id))

	case "Driver support systems":
		props = getDriverSupportSystemProperties(int(catModel.Id))

	case "Lights":
		props = getLightsProperties(int(catModel.Id))

	case "Multimedia":
		props = getMultimediaProperties(int(catModel.Id))

	case "Safety equipment":
		props = getSafetyEquipmentProperties(int(catModel.Id))

	case "Seats and steering wheel":
		props = getSeatsProperties(int(catModel.Id))

	case "Windows and mirrors":
		props = getWindowsProperties(int(catModel.Id))

	default:
		props = &([]models.Property{})
	}

	for _, prop := range *props {
		database.Model(&models.Property{}).Create(&prop)
	}
}

func getBodyProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Height",
			CategoryId:  cat,
			Description: "Height of the car",
			DataType:    "int",
			Unit:        "mm",
		}, {
			Name:        "Width",
			CategoryId:  cat,
			Description: "Width of the car",
			DataType:    "int",
			Unit:        "mm",
		}, {
			Name:        "Length",
			CategoryId:  cat,
			Description: "Length of the car",
			DataType:    "int",
			Unit:        "mm",
		},
	}

	return &props
}

func getEngineProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Displacement",
			CategoryId:  cat,
			Description: "Displacement of the car",
			DataType:    "int",
			Unit:        "cc",
		}, {
			Name:        "Engine power ",
			CategoryId:  cat,
			Description: "Engine power of the car",
			DataType:    "int",
			Unit:        "hp/kw/rpm",
		}, {
			Name:        "Maximum torque",
			CategoryId:  cat,
			Description: "Maximum torque of the car",
			DataType:    "int",
			Unit:        "nm",
		}, {
			Name:        "Fuel consumption",
			CategoryId:  cat,
			Description: "Fuel consumption (l/100km) (city/suburb/general) of the car",
			DataType:    "int",
			Unit:        "nm",
		}, {
			Name:        "Maximum speed",
			CategoryId:  cat,
			Description: "Maximum speed of the car",
			DataType:    "int",
			Unit:        "km/h",
		},
	}

	return &props
}

func getDrivetrainProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Gears",
			CategoryId:  cat,
			Description: "Gears of the car",
			DataType:    "int",
			Unit:        "count",
		},
	}

	return &props
}

func getSuspensionProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Front brakes",
			CategoryId:  cat,
			Description: "Front brakes of the car",
			DataType:    "string",
			Unit:        "",
		}, {
			Name:        "Rear brakes",
			CategoryId:  cat,
			Description: "Rear brakes of the car",
			DataType:    "string",
			Unit:        "",
		}, {
			Name:        "Tires",
			CategoryId:  cat,
			Description: "Tires of the car",
			DataType:    "string",
			Unit:        "",
		},
	}

	return &props
}

func getComfortProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Central lock",
			CategoryId:  cat,
			Description: "Central lock of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Remote lock",
			CategoryId:  cat,
			Description: "Remote lock of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Keyless engine start",
			CategoryId:  cat,
			Description: "Keyless engine start of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}

func getDriverSupportSystemProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Parking radar",
			CategoryId:  cat,
			Description: "Parking radar of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Anti-lock braking system (ABS)",
			CategoryId:  cat,
			Description: "Anti-lock braking system (ABS) of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Electronic brakeforce distribution (EBD)",
			CategoryId:  cat,
			Description: "Electronic brakeforce distribution (EBD) of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Brake Assist",
			CategoryId:  cat,
			Description: "Brake Assist of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}

func getLightsProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Low beam",
			CategoryId:  cat,
			Description: "Low beam of the car",
			DataType:    "string",
			Unit:        "",
		}, {
			Name:        "High beam",
			CategoryId:  cat,
			Description: "High beam of the car",
			DataType:    "string",
			Unit:        "",
		}, {
			Name:        "Daytime running lights",
			CategoryId:  cat,
			Description: "Daytime running lights of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Automatic headlights",
			CategoryId:  cat,
			Description: "Automatic headlights of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}

func getMultimediaProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Bluetooth",
			CategoryId:  cat,
			Description: "Bluetooth of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}

func getSafetyEquipmentProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Driver front airbag",
			CategoryId:  cat,
			Description: "Driver front airbag of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Passenger front airbag",
			CategoryId:  cat,
			Description: "Passenger front airbag of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}

func getSeatsProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Driver seat electric adjustment",
			CategoryId:  cat,
			Description: "Driver seat electric adjustment of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Passenger seat electric adjustment",
			CategoryId:  cat,
			Description: "Passenger seat electric adjustment of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}

func getWindowsProperties(cat int) *[]models.Property {
	var props []models.Property = []models.Property{
		{
			Name:        "Power windows",
			CategoryId:  cat,
			Description: "Power windows of the car",
			DataType:    "bool",
			Unit:        "",
		}, {
			Name:        "Rear wiper",
			CategoryId:  cat,
			Description: "Rear wiper of the car",
			DataType:    "bool",
			Unit:        "",
		},
	}

	return &props
}
