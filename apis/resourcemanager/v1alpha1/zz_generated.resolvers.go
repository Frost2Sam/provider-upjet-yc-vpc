/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CloudIAMBinding.
func (mg *CloudIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CloudIDRef,
		Selector:     mg.Spec.ForProvider.CloudIDSelector,
		To: reference.To{
			List:    &CloudList{},
			Managed: &Cloud{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudID")
	}
	mg.Spec.ForProvider.CloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CloudID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CloudIDRef,
		Selector:     mg.Spec.InitProvider.CloudIDSelector,
		To: reference.To{
			List:    &CloudList{},
			Managed: &Cloud{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CloudID")
	}
	mg.Spec.InitProvider.CloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CloudIAMMember.
func (mg *CloudIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CloudIDRef,
		Selector:     mg.Spec.ForProvider.CloudIDSelector,
		To: reference.To{
			List:    &CloudList{},
			Managed: &Cloud{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudID")
	}
	mg.Spec.ForProvider.CloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CloudID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CloudIDRef,
		Selector:     mg.Spec.InitProvider.CloudIDSelector,
		To: reference.To{
			List:    &CloudList{},
			Managed: &Cloud{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CloudID")
	}
	mg.Spec.InitProvider.CloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Folder.
func (mg *Folder) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CloudIDRef,
		Selector:     mg.Spec.ForProvider.CloudIDSelector,
		To: reference.To{
			List:    &CloudList{},
			Managed: &Cloud{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudID")
	}
	mg.Spec.ForProvider.CloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CloudID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CloudIDRef,
		Selector:     mg.Spec.InitProvider.CloudIDSelector,
		To: reference.To{
			List:    &CloudList{},
			Managed: &Cloud{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CloudID")
	}
	mg.Spec.InitProvider.CloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FolderIAMBinding.
func (mg *FolderIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FolderIAMMember.
func (mg *FolderIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FolderIAMPolicy.
func (mg *FolderIAMPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}