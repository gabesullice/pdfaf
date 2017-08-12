// Code generated by cdpgen. DO NOT EDIT.

// Package layertree implements the LayerTree domain.
package layertree

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the LayerTree domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the LayerTree domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the LayerTree method. Enables compositing tree inspection.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "LayerTree.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the LayerTree method. Disables compositing tree inspection.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "LayerTree.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "Disable", Err: err}
	}
	return
}

// CompositingReasons invokes the LayerTree method. Provides the reasons why the given layer was composited.
func (d *domainClient) CompositingReasons(ctx context.Context, args *CompositingReasonsArgs) (reply *CompositingReasonsReply, err error) {
	reply = new(CompositingReasonsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.compositingReasons", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.compositingReasons", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "CompositingReasons", Err: err}
	}
	return
}

// MakeSnapshot invokes the LayerTree method. Returns the layer snapshot identifier.
func (d *domainClient) MakeSnapshot(ctx context.Context, args *MakeSnapshotArgs) (reply *MakeSnapshotReply, err error) {
	reply = new(MakeSnapshotReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.makeSnapshot", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.makeSnapshot", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "MakeSnapshot", Err: err}
	}
	return
}

// LoadSnapshot invokes the LayerTree method. Returns the snapshot identifier.
func (d *domainClient) LoadSnapshot(ctx context.Context, args *LoadSnapshotArgs) (reply *LoadSnapshotReply, err error) {
	reply = new(LoadSnapshotReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.loadSnapshot", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.loadSnapshot", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "LoadSnapshot", Err: err}
	}
	return
}

// ReleaseSnapshot invokes the LayerTree method. Releases layer snapshot captured by the back-end.
func (d *domainClient) ReleaseSnapshot(ctx context.Context, args *ReleaseSnapshotArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.releaseSnapshot", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.releaseSnapshot", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "ReleaseSnapshot", Err: err}
	}
	return
}

// ProfileSnapshot invokes the LayerTree method.
func (d *domainClient) ProfileSnapshot(ctx context.Context, args *ProfileSnapshotArgs) (reply *ProfileSnapshotReply, err error) {
	reply = new(ProfileSnapshotReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.profileSnapshot", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.profileSnapshot", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "ProfileSnapshot", Err: err}
	}
	return
}

// ReplaySnapshot invokes the LayerTree method. Replays the layer snapshot and returns the resulting bitmap.
func (d *domainClient) ReplaySnapshot(ctx context.Context, args *ReplaySnapshotArgs) (reply *ReplaySnapshotReply, err error) {
	reply = new(ReplaySnapshotReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.replaySnapshot", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.replaySnapshot", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "ReplaySnapshot", Err: err}
	}
	return
}

// SnapshotCommandLog invokes the LayerTree method. Replays the layer snapshot and returns canvas log.
func (d *domainClient) SnapshotCommandLog(ctx context.Context, args *SnapshotCommandLogArgs) (reply *SnapshotCommandLogReply, err error) {
	reply = new(SnapshotCommandLogReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "LayerTree.snapshotCommandLog", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "LayerTree.snapshotCommandLog", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "LayerTree", Op: "SnapshotCommandLog", Err: err}
	}
	return
}

func (d *domainClient) LayerTreeDidChange(ctx context.Context) (DidChangeClient, error) {
	s, err := rpcc.NewStream(ctx, "LayerTree.layerTreeDidChange", d.conn)
	if err != nil {
		return nil, err
	}
	return &didChangeClient{Stream: s}, nil
}

type didChangeClient struct{ rpcc.Stream }

func (c *didChangeClient) Recv() (*DidChangeReply, error) {
	event := new(DidChangeReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "LayerTree", Op: "LayerTreeDidChange Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) LayerPainted(ctx context.Context) (LayerPaintedClient, error) {
	s, err := rpcc.NewStream(ctx, "LayerTree.layerPainted", d.conn)
	if err != nil {
		return nil, err
	}
	return &layerPaintedClient{Stream: s}, nil
}

type layerPaintedClient struct{ rpcc.Stream }

func (c *layerPaintedClient) Recv() (*LayerPaintedReply, error) {
	event := new(LayerPaintedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "LayerTree", Op: "LayerPainted Recv", Err: err}
	}
	return event, nil
}
