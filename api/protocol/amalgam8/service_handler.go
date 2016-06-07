package amalgam8

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/ant0ine/go-json-rest/rest"

	"github.com/amalgam8/registry/utils/i18n"
)

func (routes *Routes) getServiceInstances(w rest.ResponseWriter, r *rest.Request) {
	sname := r.PathParam(RouteParamServiceName)
	if sname == "" {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     "service name is required",
		}).Warn("Failed to lookup service")

		i18n.Error(r, w, http.StatusBadRequest, i18n.ErrorServiceNameMissing)
		return
	}

	catalog := routes.catalog(w, r)
	if catalog == nil {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     "catalog is nil",
		}).Errorf("Failed to lookup service %s", sname)
		// error is set in routes.catalog()
		return
	}

	if instances, err := catalog.List(sname, nil); err != nil {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     err,
		}).Errorf("Failed to lookup service %s", sname)

		i18n.Error(r, w, statusCodeFromError(err), i18n.ErrorServiceEnumeration)
		return
	} else if instances == nil || len(instances) == 0 {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     "no such service name",
		}).Warnf("Failed to lookup service %s", sname)

		i18n.Error(r, w, http.StatusNotFound, i18n.ErrorServiceNotFound)
		return
	} else {
		insts := make([]*ServiceInstance, len(instances))
		for index, si := range instances {
			inst, err := copyInstanceWithFilter(sname, si, nil)
			if err != nil {
				routes.logger.WithFields(log.Fields{
					"namespace": r.Env["REMOTE_USER"],
					"error":     err,
				}).Warnf("Failed to lookup service %s", sname)

				i18n.Error(r, w, http.StatusInternalServerError, i18n.ErrorFilterGeneric)
				return
			}
			insts[index] = inst
		}

		if err := w.WriteJson(&InstanceList{ServiceName: sname, Instances: insts}); err != nil {
			routes.logger.WithFields(log.Fields{
				"namespace": r.Env["REMOTE_USER"],
				"error":     err,
			}).Warnf("Failed to encode lookup response for %s", sname)

			i18n.Error(r, w, http.StatusInternalServerError, i18n.ErrorInternalServer)
			return
		}

		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
		}).Infof("Lookup service %s (%d)", sname, len(insts))
	}
}

func (routes *Routes) listServices(w rest.ResponseWriter, r *rest.Request) {
	catalog := routes.catalog(w, r)
	if catalog == nil {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     "catalog is nil",
		}).Error("Failed to list services")
		// error to user is already set in route.catalog()
		return
	}

	services := catalog.ListServices(nil)
	if services == nil {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     "services list is nil",
		}).Error("Failed to list services")

		i18n.Error(r, w, http.StatusInternalServerError, i18n.ErrorServiceEnumeration)
		return
	}
	listRes := &ServicesList{Services: make([]string, len(services), len(services))}

	for index, svc := range services {
		listRes.Services[index] = svc.ServiceName
	}

	err := w.WriteJson(listRes)
	if err != nil {
		routes.logger.WithFields(log.Fields{
			"namespace": r.Env["REMOTE_USER"],
			"error":     err,
		}).Warn("Failed to encode services list")

		i18n.Error(r, w, http.StatusInternalServerError, i18n.ErrorEncoding)
		return
	}

	routes.logger.WithFields(log.Fields{
		"namespace": r.Env["REMOTE_USER"],
	}).Infof("List services (%d)", len(listRes.Services))
}
