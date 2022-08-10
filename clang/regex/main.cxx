#include <iostream>
#include <iterator>
#include <string>
#include <regex>

int main()
{
    std::string stat_name = "cluster.xds-grpc.upstream_rq_200";

    std::regex self_regex("REGULAR EXPRESSIONS",
                          std::regex_constants::ECMAScript | std::regex_constants::icase);
    if (std::regex_search(stat_name, self_regex))
    {
        std::cout << "Text contains the phrase 'regular expressions'\n";
    }

    std::regex word_regex("(response_code=\\.=(.+?);\\.;)|_rq(_(\\d{3}))");
    std::match_results<std::string::iterator> match;
    if (std::regex_search<std::string::iterator>(stat_name.begin(), stat_name.end(), match, word_regex) && match.size() > 1)
    {
        std::cout << "match size: " << match.size() << std::endl;
        for (auto iter = match.begin(); iter != match.end(); iter++)
        {
            std::cout << iter->str() << std::endl;
        }
    }

    std::regex word_regex2("_rq(_(\\d{3}))|(response_code=\\.=(.+?);\\.;)");
    std::match_results<std::string::iterator> match2;
    if (std::regex_search<std::string::iterator>(stat_name.begin(), stat_name.end(), match2, word_regex2) && match2.size() > 1)
    {
        std::cout << "match size: " << match2.size() << std::endl;
        for (auto iter = match2.begin(); iter != match2.end(); iter++)
        {
            std::cout << iter->str() << std::endl;
        }
    }
}