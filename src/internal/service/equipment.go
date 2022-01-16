package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type EquipmentService struct {
	equipment repository.EquipmentRepo
}

func NewEquipmentService(eq repository.EquipmentRepo) *EquipmentService {
	return &EquipmentService{equipment: eq}
}

func (e *EquipmentService) GetAll() (equipments *[]domain.Equipment, err error) {
	equipments, err = e.equipment.GetAll()
	return
	//if err != nil {
	//	http.Error(w, http.StatusText(500), 500)
	//	return
	//}
	//enc := json.NewEncoder(w)
	//if err := enc.Encode(equipments); err != nil {
	//	http.Error(w, http.StatusText(500), 500)
	//	return
	//}
}

func (e *EquipmentService) CreateEquipment(in *domain.Equipment) (equipment *domain.Equipment, err error) {
	if err := e.equipment.CreateEquipment(in); err != nil {
		return nil, err
	}
	return in, nil

	//data, err := io.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, http.StatusText(500), 500)
	//	return
	//}
	//eqt := new(domain.Equipment)
	//if err = json.Unmarshal(data, eqt); err != nil {
	//	http.Error(w, http.StatusText(422), 422)
	//	logrus.Debugf("Couldn't unmarshal: %s", err)
	//	return
	//}
	//if err = e.equipment.createEquipment(eqt); err != nil {
	//	http.Error(w, http.StatusText(500), 500)
	//	logrus.Debugf("Couldn't create %#v: %s", eqt, err)
	//}
}
