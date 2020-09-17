unit auto_linkup_GUI_unit;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ComCtrls, StdCtrls;

type

  { TAutoLinkupGUI }

  TAutoLinkupGUI = class(TForm)
    ReleaseRateLabel: TLabel;
    ReleaseRateBar: TTrackBar;
    procedure ReleaseRateBarChange(Sender: TObject);
  private

  public

  end;

var
  AutoLinkupGUI: TAutoLinkupGUI;

implementation

{$R *.lfm}

{ TAutoLinkupGUI }

procedure TAutoLinkupGUI.ReleaseRateBarChange(Sender: TObject);
begin

end;

end.

