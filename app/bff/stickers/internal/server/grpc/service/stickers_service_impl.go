package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/stickers/internal/core"
)

// MessagesGetStickers
func (s *Service) MessagesGetStickers(ctx context.Context, request *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetAllStickers
func (s *Service) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetAllStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetAllStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetAllStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetStickerSet
func (s *Service) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetStickerSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetStickerSet - reply: {%s}", r)
	return r, err
}
// MessagesInstallStickerSet
func (s *Service) MessagesInstallStickerSet(ctx context.Context, request *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesInstallStickerSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesInstallStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesInstallStickerSet - reply: {%s}", r)
	return r, err
}
// MessagesUninstallStickerSet
func (s *Service) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesUninstallStickerSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesUninstallStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesUninstallStickerSet - reply: {%s}", r)
	return r, err
}
// MessagesReorderStickerSets
func (s *Service) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReorderStickerSets - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReorderStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReorderStickerSets - reply: {%s}", r)
	return r, err
}
// MessagesGetFeaturedStickers
func (s *Service) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetFeaturedStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetFeaturedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetFeaturedStickers - reply: {%s}", r)
	return r, err
}
// MessagesReadFeaturedStickers
func (s *Service) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReadFeaturedStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReadFeaturedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReadFeaturedStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetRecentStickers
func (s *Service) MessagesGetRecentStickers(ctx context.Context, request *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetRecentStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetRecentStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetRecentStickers - reply: {%s}", r)
	return r, err
}
// MessagesSaveRecentSticker
func (s *Service) MessagesSaveRecentSticker(ctx context.Context, request *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSaveRecentSticker - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSaveRecentSticker(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSaveRecentSticker - reply: {%s}", r)
	return r, err
}
// MessagesClearRecentStickers
func (s *Service) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesClearRecentStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesClearRecentStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesClearRecentStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetArchivedStickers
func (s *Service) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetArchivedStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetArchivedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetArchivedStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetMaskStickers
func (s *Service) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetMaskStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetMaskStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetMaskStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetAttachedStickers
func (s *Service) MessagesGetAttachedStickers(ctx context.Context, request *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetAttachedStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetAttachedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetAttachedStickers - reply: {%s}", r)
	return r, err
}
// MessagesGetFavedStickers
func (s *Service) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetFavedStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetFavedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetFavedStickers - reply: {%s}", r)
	return r, err
}
// MessagesFaveSticker
func (s *Service) MessagesFaveSticker(ctx context.Context, request *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesFaveSticker - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesFaveSticker(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesFaveSticker - reply: {%s}", r)
	return r, err
}
// MessagesSearchStickerSets
func (s *Service) MessagesSearchStickerSets(ctx context.Context, request *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSearchStickerSets - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSearchStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSearchStickerSets - reply: {%s}", r)
	return r, err
}
// MessagesToggleStickerSets
func (s *Service) MessagesToggleStickerSets(ctx context.Context, request *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesToggleStickerSets - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesToggleStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesToggleStickerSets - reply: {%s}", r)
	return r, err
}
// MessagesGetOldFeaturedStickers
func (s *Service) MessagesGetOldFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetOldFeaturedStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetOldFeaturedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetOldFeaturedStickers - reply: {%s}", r)
	return r, err
}
// MessagesSearchEmojiStickerSets
func (s *Service) MessagesSearchEmojiStickerSets(ctx context.Context, request *mtproto.TLMessagesSearchEmojiStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSearchEmojiStickerSets - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSearchEmojiStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSearchEmojiStickerSets - reply: {%s}", r)
	return r, err
}
// MessagesGetMyStickers
func (s *Service) MessagesGetMyStickers(ctx context.Context, request *mtproto.TLMessagesGetMyStickers) (*mtproto.Messages_MyStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetMyStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetMyStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetMyStickers - reply: {%s}", r)
	return r, err
}
// MessagesSearchStickers
func (s *Service) MessagesSearchStickers(ctx context.Context, request *mtproto.TLMessagesSearchStickers) (*mtproto.Messages_FoundStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSearchStickers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSearchStickers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSearchStickers - reply: {%s}", r)
	return r, err
}
// StickersCreateStickerSet
func (s *Service) StickersCreateStickerSet(ctx context.Context, request *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersCreateStickerSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersCreateStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersCreateStickerSet - reply: {%s}", r)
	return r, err
}
// StickersRemoveStickerFromSet
func (s *Service) StickersRemoveStickerFromSet(ctx context.Context, request *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersRemoveStickerFromSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersRemoveStickerFromSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersRemoveStickerFromSet - reply: {%s}", r)
	return r, err
}
// StickersChangeStickerPosition
func (s *Service) StickersChangeStickerPosition(ctx context.Context, request *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersChangeStickerPosition - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersChangeStickerPosition(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersChangeStickerPosition - reply: {%s}", r)
	return r, err
}
// StickersAddStickerToSet
func (s *Service) StickersAddStickerToSet(ctx context.Context, request *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersAddStickerToSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersAddStickerToSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersAddStickerToSet - reply: {%s}", r)
	return r, err
}
// StickersSetStickerSetThumb
func (s *Service) StickersSetStickerSetThumb(ctx context.Context, request *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersSetStickerSetThumb - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersSetStickerSetThumb(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersSetStickerSetThumb - reply: {%s}", r)
	return r, err
}
// StickersCheckShortName
func (s *Service) StickersCheckShortName(ctx context.Context, request *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersCheckShortName - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersCheckShortName(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersCheckShortName - reply: {%s}", r)
	return r, err
}
// StickersSuggestShortName
func (s *Service) StickersSuggestShortName(ctx context.Context, request *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersSuggestShortName - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersSuggestShortName(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersSuggestShortName - reply: {%s}", r)
	return r, err
}
// StickersChangeSticker
func (s *Service) StickersChangeSticker(ctx context.Context, request *mtproto.TLStickersChangeSticker) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersChangeSticker - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersChangeSticker(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersChangeSticker - reply: {%s}", r)
	return r, err
}
// StickersRenameStickerSet
func (s *Service) StickersRenameStickerSet(ctx context.Context, request *mtproto.TLStickersRenameStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersRenameStickerSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersRenameStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersRenameStickerSet - reply: {%s}", r)
	return r, err
}
// StickersDeleteStickerSet
func (s *Service) StickersDeleteStickerSet(ctx context.Context, request *mtproto.TLStickersDeleteStickerSet) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersDeleteStickerSet - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersDeleteStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersDeleteStickerSet - reply: {%s}", r)
	return r, err
}
// StickersReplaceSticker
func (s *Service) StickersReplaceSticker(ctx context.Context, request *mtproto.TLStickersReplaceSticker) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StickersReplaceSticker - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StickersReplaceSticker(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StickersReplaceSticker - reply: {%s}", r)
	return r, err
}
