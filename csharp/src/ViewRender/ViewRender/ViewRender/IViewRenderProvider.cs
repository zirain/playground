using System.Threading.Tasks;

namespace ViewRender
{
    public interface IViewRenderProvider
    {
        Task<string> RenderToStringAsync(string viewName, object model, bool isMainPage = false);
    }
}
